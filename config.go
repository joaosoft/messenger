package messenger

import (
	"fmt"

	"github.com/joaosoft/socket"

	"github.com/joaosoft/dbr"
	"github.com/joaosoft/manager"
	migration "github.com/joaosoft/migration/services"
)

// AppConfig ...
type AppConfig struct {
	Messenger *MessengerConfig `json:"messenger"`
}

// MessengerConfig ...
type MessengerConfig struct {
	Host              string                     `json:"host"`
	Socket            *socket.SocketConfig       `json:"socket"`
	Dbr               *dbr.DbrConfig             `json:"dbr"`
	TokenKey          string                     `json:"token_key"`
	ExpirationMinutes int64                      `json:"expiration_minutes"`
	Migration         *migration.MigrationConfig `json:"migration"`
	Log               struct {
		Level string `json:"level"`
	} `json:"log"`
}

// newConfig ...
func NewConfig() (*AppConfig, manager.IConfig, error) {
	appConfig := &AppConfig{}
	simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", getEnv()), appConfig)

	return appConfig, simpleConfig, err
}
