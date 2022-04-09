package main

import (
	"github.com/eaardal/workwork/src/commands"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "WorkWork",
		Usage: "A simple dictionary for listing and opening URLs for common software development concerns",
		Commands: []*cli.Command{
			commands.InitCommand,
			commands.LSCommand,
			commands.GoToCommand,
			commands.SetCommand,
			commands.GetCommand,
			commands.RMCommand,
			commands.ExecCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
