package web

import (
	"github.com/labstack/echo"
	"github.com/lcaballero/itemize/internal/svc/da"
	"net/http"
)


type UsersEndpoint struct {
	data *da.DataStore
	Render *Render
}

func NewUsersEndpoint(data *da.DataStore) *UsersEndpoint {
	return &UsersEndpoint{
		Render: NewRender(DefaultBaseTemplaDir),
		data: data,
	}
}

func (u *UsersEndpoint) Read(c *echo.Context) error {
	html, err := u.Render.Render("list_users.tmpl", u.data)
	if err != nil {
		return err
	} else {
		return c.HTML(http.StatusOK, html)
	}
}

func (u *UsersEndpoint) New(c *echo.Context) error {
	html, err := u.Render.Render("new_user.tmpl", nil)
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

	u.data.AddUser(user)

	html, err := u.Render.Render("read_user.tmpl", user)
	if err != nil {
		return err
	} else {
		return c.HTML(http.StatusOK, html)
	}
}


