package web

import (
	"net/http"

	"log"

	"github.com/labstack/echo"
	"github.com/lcaballero/itemize/internal/svc/da"
)

type UsersEndpoint struct {
	client *da.AccessClient
	Render *Render
}

func NewUsersEndpoint(client *da.AccessClient) *UsersEndpoint {
	return &UsersEndpoint{
		Render: &Render{},
		client: client,
	}
}

func (u *UsersEndpoint) Read(c *echo.Context) error {
	users := u.client.Users()
	html, err := u.Render.Render(users, WebDir("list_users.html"))
	if err != nil {
		return err
	} else {
		return c.HTML(http.StatusOK, html)
	}
}

func (u *UsersEndpoint) New(c *echo.Context) error {
	html, err := u.Render.Render(nil, WebDir("new_user.html"))
	if err != nil {
		return err
	} else {
		return c.HTML(http.StatusOK, html)
	}
}

func (u *UsersEndpoint) User(c *echo.Context) error {
	id := c.Param("id")

	user, err := u.client.FindUser(id)
	if err != nil {
		log.Println(err)
		return err
	}

	html, err := u.Render.Render(user, WebDir("user.html"))
	if err != nil {
		return err
	} else {
		return c.HTML(http.StatusOK, html)
	}
}

func (u *UsersEndpoint) Edit(c *echo.Context) error {
	id := c.Param("id")

	user, err := u.client.FindUser(id)
	if err != nil {
		log.Println(err)
		return err
	}

	html, err := u.Render.Render(user, WebDir("edit_user.html"))
	if err != nil {
		return err
	} else {
		return c.HTML(http.StatusOK, html)
	}
}

func (u *UsersEndpoint) Update(c *echo.Context) error {
	user := da.User{
		Username:  c.Form("Username"),
		FirstName: c.Form("FirstName"),
		LastName:  c.Form("LastName"),
		Email:     c.Form("Email"),
		Id:        c.Param("id"),
	}

	log.Println("id", user.Id)
	user, err := u.client.UpdateUser(user)
	if err != nil {
		log.Println(err)
		return err
	}

	html, err := u.Render.Render(user, WebDir("user.html"))
	if err != nil {
		return err
	} else {
		return c.HTML(http.StatusOK, html)
	}
}

func (u *UsersEndpoint) Create(c *echo.Context) error {
	user := da.NewUser()
	user.Username = c.Form("Username")
	user.FirstName = c.Form("FirstName")
	user.LastName = c.Form("LastName")
	user.Email = c.Form("Email")

	u.client.AddUser(user)

	html, err := u.Render.Render(user, WebDir("user.html"))
	if err != nil {
		return err
	} else {
		return c.HTML(http.StatusOK, html)
	}
}
