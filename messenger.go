package messenger

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
	migration "github.com/joaosoft/migration/services"
	"github.com/joaosoft/socket"
)

type Messenger struct {
	config        *MessengerConfig
	socketServer  *socket.Server
	socketClient  *socket.Client
	isLogExternal bool
	pm            *manager.Manager
	logger        logger.ILogger
	mux           sync.Mutex
}

// NewMessenger ...
func NewMessenger(options ...SessionOption) (*Messenger, error) {
	config, simpleConfig, err := NewConfig()

	service := &Messenger{
		pm:     manager.NewManager(manager.WithRunInBackground(true)),
		logger: logger.NewLogDefault("messenger", logger.WarnLevel),
		config: config.Messenger,
	}

	if err != nil {
		service.logger.Error(err.Error())
	} else if config.Messenger != nil {
		service.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.Messenger.Log.Level)
		service.logger.Debugf("setting log level to %s", level)
		service.logger.Reconfigure(logger.WithLevel(level))
	} else {
		config.Messenger = &MessengerConfig{
			Host: defaultURL,
		}
	}

	service.Reconfigure(options...)

	service.socketServer, err = socket.NewServer(socket.WithServerManager(service.pm), socket.WithServerConfiguration(config.Messenger.Socket.Server))
	if err != nil {
		panic(err)
	}

	service.socketClient, err = socket.NewClient(socket.WithClientManager(service.pm), socket.WithClientConfiguration(config.Messenger.Socket.Client))
	if err != nil {
		panic(err)
	}

	if service.isLogExternal {
		service.pm.Reconfigure(manager.WithLogger(logger.Instance))
	}

	// execute migrations
	migrationService, err := migration.NewCmdService(migration.WithCmdConfiguration(service.config.Migration))
	if err != nil {
		return nil, err
	}

	if _, err := migrationService.Execute(migration.OptionUp, 0, migration.ExecutorModeDatabase); err != nil {
		return nil, err
	}

	web := service.pm.NewSimpleWebServer(config.Messenger.Host)

	storage, err := NewStoragePostgres(config.Messenger)
	if err != nil {
		return nil, err
	}

	interactor := NewInteractor(config.Messenger, storage, service.socketClient)

	controller := NewController(config.Messenger, interactor)
	controller.RegisterRoutes(web)

	service.pm.AddWeb("api_web_messenger", web)

	return service, nil
}

// Start ...
func (m *Messenger) Start() error {
	err := m.pm.Start()

	if err == nil {
		m.wait()
	}

	return err
}

// Stop ...
func (m *Messenger) Stop() error {
	return m.pm.Stop()
}

// Forget ...
func (c *Messenger) wait() {
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	select {
	case <-termChan:
		c.Stop()
		c.logger.Infof("received term signal")
	}
}
