package main

import (
	"os"

	"github.com/lcaballero/itemize/internal/cli"
)

func main() {
	cli.NewCli().Run(os.Args)
}
