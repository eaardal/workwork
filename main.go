package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			initCmd,
			ls,
			set,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
