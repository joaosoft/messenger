package messenger

import (
	"time"

	"github.com/joaosoft/web"
)

type ErrorResponse struct {
	Code    web.Status `json:"code,omitempty"`
	Message string     `json:"message,omitempty"`
	Cause   string     `json:"cause,omitempty"`
}

type SendMessageRequest struct {
	From string `json:"from" validate:"notzero"`
	To   string `json:"to" validate:"notzero"`
	Body struct {
		Message string `json:"message"`
	} `json:"body"`
}

type SendMessageResponse struct {
	Success bool `json:"success"`
}

type GetMessagesRequest struct {
	User string `json:"user" validate:"notzero"`
}

type GetMessagesResponse []*Message

type Message struct {
	IdMessage string    `json:"id_message" db:"id_message"`
	From      string    `json:"from" db:"from"`
	To        string    `json:"to" db:"to"`
	Message   string    `json:"message" db:"message"`
	CreatedAt time.Time `json:"created_at" db.read:"created_at"`
}
