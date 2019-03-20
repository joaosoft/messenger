package messenger

import (
	logger "github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

// SessionOption ...
type SessionOption func(client *Messenger)

// Reconfigure ...
func (session *Messenger) Reconfigure(options ...SessionOption) {
	for _, option := range options {
		option(session)
	}
}

// WithConfiguration ...
func WithConfiguration(config *MessengerConfig) SessionOption {
	return func(session *Messenger) {
		session.config = config
	}
}

// WithLogger ...
func WithLogger(logger logger.ILogger) SessionOption {
	return func(session *Messenger) {
		log = logger
		session.isLogExternal = true
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) SessionOption {
	return func(session *Messenger) {
		log.SetLevel(level)
	}
}

// WithManager ...
func WithManager(mgr *manager.Manager) SessionOption {
	return func(session *Messenger) {
		session.pm = mgr
	}
}
