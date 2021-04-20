package commands

import (
	"github.com/urfave/cli/v2"
	"workwork/src/gui"
	"workwork/src/ww"
)

var LSCommand = &cli.Command{
	Name:      "ls",
	Usage:     "List all the registered URLs",
	UsageText: "Run like this: `ww ls`",
	Action: func(c *cli.Context) error {
		wwFile, err := ww.ReadWorkWorkFile()
		if err != nil {
			return err
		}

		ui := gui.NewUserInterface()

		for key, value := range wwFile.Urls {
			ui.Write("%s\t%s\t", gui.FgHiWhite(key), value)
		}

		for _, env := range wwFile.Environments {
			ui.Write("\n%s", gui.FgHiGreen(env.Name))

			for key, value := range env.Urls {
				ui.Write("%s\t%s\t", gui.FgHiWhite(key), value)
			}
		}

		ui.MustFlush()
		return nil
	},
}
