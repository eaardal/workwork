package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "WorkWork",
		Usage: "A simple dictionary for listing and opening URLs for common software development concerns",
		Commands: []*cli.Command{
			initCommand,
			lsCommand,
			gotoCommand,
			setCommand,
			rmCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
