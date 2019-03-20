package messenger

import (
	"fmt"

	"github.com/joaosoft/dbr"
)

type StoragePostgres struct {
	config *MessengerConfig
	db     *dbr.Dbr
}

func NewStoragePostgres(config *MessengerConfig) (*StoragePostgres, error) {
	dbr, err := dbr.New(dbr.WithConfiguration(config.Dbr))
	if err != nil {
		return nil, err
	}

	return &StoragePostgres{
		config: config,
		db:     dbr,
	}, nil
}

func (storage *StoragePostgres) SaveMessage(message *Message) error {
	str, _ := storage.db.
		Insert().
		Into("message").
		Record(message).
		Build()
	fmt.Println(str)
	_, err := storage.db.
		Insert().
		Into("message").
		Record(message).
		Exec()

	if err != nil {
		return err
	}

	return nil
}

func (storage *StoragePostgres) GetMessages(user string) ([]*Message, error) {
	var messages []*Message

	_, err := storage.db.
		Select("*").
		From("message").
		Where(`"from" = ?`, user).
		WhereOr(`"to" = ?`, user).
		Limit(10).
		OrderDesc("created_at").
		Load(&messages)

	if err != nil {
		return nil, err
	}

	return messages, nil
}
