package web

import (
	"net/http"

	"log"

	"github.com/labstack/echo"
	"github.com/lcaballero/itemize/internal/svc/da"
)

type ListItems struct {
	Render *Render
	client *da.AccessClient
}

func NewListEndpoint(client *da.AccessClient) *ListItems {
	return &ListItems{
		Render: &Render{},
		client: client,
	}
}

func (m *ListItems) Read(c *echo.Context) error {
	items := m.client.Items()
	log.Println("items", items)
	html, err := m.Render.Render(items, WebDir("list_items.html"))
	if err != nil {
		return err
	} else {
		return c.HTML(http.StatusOK, html)
	}
}

func (m *ListItems) Item(c *echo.Context) error {
	id := c.Param("id")

	item, err := m.client.FindItem(id)
	if err != nil {
		log.Println(err)
		return err
	}

	html, err := m.Render.Render(item, WebDir("item.html"))
	if err != nil {
		return err
	} else {
		return c.HTML(http.StatusOK, html)
	}
}

func (m *ListItems) Edit(c *echo.Context) error {
	id := c.Param("id")

	item, err := m.client.FindItem(id)
	if err != nil {
		log.Println(err)
		return err
	}

	html, err := m.Render.Render(item, WebDir("edit_item.html"))
	if err != nil {
		return err
	} else {
		return c.HTML(http.StatusOK, html)
	}
}

func (m *ListItems) New(c *echo.Context) error {
	html, err := m.Render.Render(nil, WebDir("new_item.html"))
	if err != nil {
		return err
	} else {
		return c.HTML(http.StatusOK, html)
	}
}

func (m *ListItems) Update(c *echo.Context) error {
	item := da.Item{
		Title:   c.Form("Title"),
		Summary: c.Form("Summary"),
	}
	item, err := m.client.UpdateItem(c.Param("id"), item)
	if err != nil {
		return err
	}

	html, err := m.Render.Render(item, WebDir("item.html"))
	if err != nil {
		return err
	} else {
		return c.HTML(http.StatusOK, html)
	}
}

func (m *ListItems) Create(c *echo.Context) error {
	item := da.NewItem()
	item.Title = c.Form("Title")
	item.Summary = c.Form("Summary")

	m.client.AddItem(*item)

	html, err := m.Render.Render(item, WebDir("item.html"))
	if err != nil {
		return err
	} else {
		return c.HTML(http.StatusOK, html)
	}
}
