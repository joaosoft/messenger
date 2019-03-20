package messenger

import (
	"encoding/json"

	"github.com/joaosoft/socket"
)

type IStorage interface {
	SaveMessage(message *Message) error
	GetMessages(user string) ([]*Message, error)
}

type Interactor struct {
	config  *MessengerConfig
	socket  *socket.Client
	storage IStorage
}

func NewInteractor(config *MessengerConfig, storageDB IStorage, socket *socket.Client) *Interactor {
	return &Interactor{
		config:  config,
		socket:  socket,
		storage: storageDB,
	}
}

func (i *Interactor) SaveMessage(message *Message) error {
	log.WithFields(map[string]interface{}{"method": "SaveMessage"})
	log.Infof("adding new message [from user: %s, to user: %s] %s", message.From, message.To)

	err := i.storage.SaveMessage(message)
	if err != nil {
		log.WithFields(map[string]interface{}{"error": err.Error()}).
			Errorf("error adding new message [from: %s, to: %s] %s", message.From, message.To, err).ToError()
		return err
	}

	// dispatch message to the other user queue
	msg, err := json.Marshal(message)
	if err != nil {
		log.WithFields(map[string]interface{}{"error": err.Error()}).
			Errorf("error adding new message [from: %s, to: %s] %s", message.From, message.To, err).ToError()
		return err
	}

	if err := i.socket.Publish(message.To, "messenger::in-box", msg); err != nil {
		panic(err)
	}

	return nil
}

func (i *Interactor) GetMessages(user string) ([]*Message, error) {
	log.WithFields(map[string]interface{}{"method": "GetMessages"})
	log.Infof("getting messages of [user: %s]", user)

	messages, err := i.storage.GetMessages(user)
	if err != nil {
		log.WithFields(map[string]interface{}{"error": err.Error()}).
			Errorf("error getting messages of [user: %s] %s", user, err).ToError()
		return nil, err
	}

	return messages, nil
}
