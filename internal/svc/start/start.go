package start

import (
	"syscall"

	cmd "github.com/codegangsta/cli"
	"github.com/lcaballero/hitman"
	"github.com/lcaballero/itemize/internal/svc/da"
	"github.com/lcaballero/itemize/internal/svc/web"
	"github.com/vrecan/death"
)

func Start(cli *cmd.Context) {
	dbname := cli.String("filename")

	client, err := da.NewAccessClient(dbname)
	targets := hitman.NewTargets()
	targets.AddOrPanic(client, err)
	targets.AddOrPanic(web.NewWebServer(client))

	death.NewDeath(syscall.SIGTERM, syscall.SIGINT).WaitForDeath(targets)
}
