package messenger

import (
	"github.com/joaosoft/errors"
	"github.com/joaosoft/web"
)

var (
	ErrorNotFound = errors.New(errors.LevelError, int(web.StatusNotFound), "not found")
)
