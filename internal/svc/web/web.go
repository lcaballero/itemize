package web

import (
	"log"

	"path/filepath"

	"github.com/labstack/echo"
	"github.com/lcaballero/hitman"
	"github.com/lcaballero/itemize/internal/svc/da"
)

const WebRoot = ".www/dest/_site/html"
const DefaultTemplateBaseDir = ".web/src/templates/tmpl"
const DefaultPartialsBaseDir = ".web/src/templates/partials"
const DefaultLayoutBaseDir = ".web/src/templates/layout"

func WebDir(elem ...string) string {
	elems := []string{WebRoot}
	elems = append(elems, elem...)
	return filepath.Join(elems...)
}

type WebServer struct {
	client *da.AccessClient
}

func NewWebServer(access *da.AccessClient) (*WebServer, error) {
	w := &WebServer{
		client: access,
	}
	return w, nil
}

func (w *WebServer) Start() hitman.KillChannel {
	done := hitman.NewKillChannel()
	go func() {
		go w.run(w.client)
		for {
			select {
			case cleaner := <-done:
				log.Println("Stopping web server")
				cleaner.WaitGroup.Done()
				return
			}
		}
	}()
	return done
}

func (w *WebServer) run(client *da.AccessClient) {
	log.Println("Starting web server")

	e := echo.New()
	e.Get("/users", NewUsersEndpoint(client).Read)
	e.Get("/new/user", NewUsersEndpoint(client).New)
	e.Post("/add/user", NewUsersEndpoint(client).Create)
	e.Get("/user/:id", NewUsersEndpoint(client).User)
	e.Get("/edit/user/:id", NewUsersEndpoint(client).Edit)
	e.Post("/update/user/:id", NewUsersEndpoint(client).Update)

	e.Get("/items", NewListEndpoint(client).Read)
	e.Get("/new/item", NewListEndpoint(client).New)
	e.Post("/add/item", NewListEndpoint(client).Create)
	e.Get("/item/:id", NewListEndpoint(client).Item)
	e.Get("/edit/item/:id", NewListEndpoint(client).Edit)
	e.Post("/update/item/:id", NewListEndpoint(client).Update)

	e.Index(".www/dest/_site/html/index.html")
	e.Static("/css", ".www/dest/_site/css")

	log.Println("Web started :2222")
	e.Run(":2222")
}
