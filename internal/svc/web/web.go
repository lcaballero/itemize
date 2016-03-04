package web

import (
	"log"

	"github.com/labstack/echo"
	"github.com/lcaballero/hitman"
	"github.com/lcaballero/itemize/internal/svc/da"
)

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
	e.Get("/", NewHelloWorld(data).Hello)
	e.Run(":2222")
}
