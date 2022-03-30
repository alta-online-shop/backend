package response

import (
	"github.com/ngamux/ctx"
	"github.com/ngamux/ngamux"
)

func Success(c *ctx.Context, status int, data interface{}) error {
	return c.JSON(status, ngamux.Map{
		"data": data,
	})
}

func Fail(c *ctx.Context, status int, err string) error {
	return c.JSON(status, ngamux.Map{
		"error": err,
	})
}
