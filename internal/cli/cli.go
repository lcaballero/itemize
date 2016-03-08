package cli

import (
	cmd "github.com/codegangsta/cli"
	"github.com/lcaballero/itemize/internal/svc/da"
	"github.com/lcaballero/itemize/internal/svc/start"
)

func NewCli() *cmd.App {
	app := cmd.NewApp()
	app.Name = "itemize"
	app.Version = "0.0.1"
	app.Usage = "Runs the itemize web app server."
	app.Commands = []cmd.Command{
		runServerCommand(),
	}
	return app
}

func runServerCommand() cmd.Command {
	return cmd.Command{
		Name:   "web",
		Action: start.Start,
		Flags: []cmd.Flag{
			cmd.StringFlag{
				Name:  "filename",
				Value: da.DefaultDbName,
				Usage: "Specify the name of the database file",
			},
		},
	}
}
