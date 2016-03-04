package start

import (
	"syscall"

	"github.com/lcaballero/hitman"
	"github.com/lcaballero/itemize/internal/svc/da"
	"github.com/lcaballero/itemize/internal/svc/web"
	"github.com/vrecan/death"
)

func Start() {
	dbname := ""
	data := make(chan *da.DataStore)

	targets := NewTargets()
	targets.AddOrPanic(web.NewWebServer(data))
	targets.AddOrPanic(da.NewDataWriter(dbname, data))

	death.NewDeath(syscall.SIGTERM, syscall.SIGINT).WaitForDeath(targets)
}

type Targets struct {
	hitman.Targets
}

func NewTargets() *Targets {
	return &Targets{
		Targets: hitman.NewTargets(),
	}
}
