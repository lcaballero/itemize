package web

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lcaballero/itemize/internal/svc/da"
)

type HelloWorld struct {
	data *da.DataStore
}

func NewHelloWorld(data *da.DataStore) *HelloWorld {
	return &HelloWorld{}
}

func (h *HelloWorld) Hello(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
