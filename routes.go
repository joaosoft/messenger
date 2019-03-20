package messenger

import (
	"net/http"

	"github.com/joaosoft/manager"
)

func (c *Controller) RegisterRoutes(web manager.IWeb) error {
	return web.AddRoutes(
		manager.NewRoute(http.MethodGet, "/api/v1/messenger/messages", c.GetMessagesHandler),
		manager.NewRoute(http.MethodPost, "/api/v1/messenger/message/users/:id", c.SendMessageHandler),
	)
}
