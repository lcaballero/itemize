package web

import (
	"log"

	"github.com/labstack/echo"
	"github.com/lcaballero/hitman"
	"github.com/lcaballero/itemize/internal/svc/da"
)

const DefaultBaseTemplaDir = ".web/src/tmpl"

type WebServer struct {
	data chan *da.DataStore
}

func NewWebServer(data chan *da.DataStore) (*WebServer, error) {
	w := &WebServer{
		data: data,
	}
	return w, nil
}

func (w *WebServer) Start() hitman.KillChannel {
	done := hitman.NewKillChannel()

	log.Println("Awaiting DataStore startup")
	go func() {
		for {
			select {
			case cleaner := <-done:
				log.Println("Stopping web server")
				cleaner.WaitGroup.Done()
				return
			case data := <-w.data:
				go w.run(data)
			}
		}
	}()

	return done
}

func (w *WebServer) run(data *da.DataStore) {
	log.Println("Starting web server")
	e := echo.New()
	e.Get("/items", NewListEndpoint(data).Read)
	e.Get("/users", NewUsersEndpoint(data).Read)
	e.Get("/new/user", NewUsersEndpoint(data).New)
	e.Post("/add/user", NewUsersEndpoint(data).Create)


	e.Index(".web/src/html/index.html")
	e.Static("/js", ".web/src/js")
	e.Static("/css", ".web/src/css")
	e.Static("/html", ".web/src/html")
	e.Run(":2222")
}
