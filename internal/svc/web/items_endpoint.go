package web

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/lcaballero/itemize/internal/svc/da"
)

type ListItems struct {
	Render *Render
	data *da.DataStore
}

func NewListEndpoint(data *da.DataStore) *ListItems {
	return &ListItems{
		Render: NewRender(DefaultBaseTemplaDir),
		data: data,
	}
}

func (m *ListItems) Read(c *echo.Context) error {
	html, err := m.Render.Render("list_items.tmpl", m.data)
	if err != nil {
		return err
	} else {
		return c.HTML(http.StatusOK, html)
	}
}
