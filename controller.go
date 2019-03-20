package messenger

import (
	"encoding/json"

	"github.com/joaosoft/validator"
	"github.com/joaosoft/web"
)

type Controller struct {
	config     *MessengerConfig
	interactor *Interactor
}

func NewController(config *MessengerConfig, interactor *Interactor) *Controller {
	return &Controller{
		config:     config,
		interactor: interactor,
	}
}

func (c *Controller) SendMessageHandler(ctx *web.Context) error {
	request := &SendMessageRequest{
		From: ctx.Request.GetHeader("user"),
		To:   ctx.Request.GetUrlParam("id"),
	}

	err := json.Unmarshal(ctx.Request.Body, &request.Body)
	if err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	err = c.interactor.SaveMessage(&Message{
		IdMessage: genUI(),
		From:      request.From,
		To:        request.To,
		Message:   request.Body.Message,
	})

	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, ErrorResponse{Code: web.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.Response.JSON(web.StatusOK, &SendMessageResponse{Success: true})
}

func (c *Controller) GetMessagesHandler(ctx *web.Context) error {
	request := &GetMessagesRequest{
		User: ctx.Request.GetHeader("user"),
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	response, err := c.interactor.GetMessages(request.User)

	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, ErrorResponse{Code: web.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.Response.JSON(web.StatusOK, GetMessagesResponse(response))
}
