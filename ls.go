package main

import (
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

var ls = &cli.Command{
	Name:  "ls",
	Usage: "List all the registered URLs",
	Action: func(c *cli.Context) error {
		ww, err := readWorkWorkFile()
		if err != nil {
			return err
		}

		hiWhite := color.New(color.FgHiWhite).SprintFunc()
		printer := newUserInterface()

		for key, value := range ww.Urls {
			printer.write("%s\t%s", hiWhite(key), value)
		}

		printer.mustFlush()
		return nil
	},
}
