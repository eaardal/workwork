package commands

import (
	"github.com/urfave/cli/v2"
	"workwork/src/gui"
	"workwork/src/ww"
)

var LSCommand = &cli.Command{
	Name:      "ls",
	Usage:     "List all the registered URLs",
	UsageText: "Example: `ww ls`",
	Action: func(c *cli.Context) error {
		wwFile, err := ww.ReadWorkWorkFile()
		if err != nil {
			return err
		}

		ui := gui.NewUserInterface()

		ui.Write("%s", gui.FgHiGreen("global"))

		for key, value := range wwFile.GlobalUrls {
			ui.Write("%s\t%s\t", gui.FgHiWhite(key), value)
		}

		for _, env := range wwFile.Environments {
			ui.Write("\n%s", gui.FgHiGreen(env.Name))

			for key, value := range env.EnvironmentUrls {
				ui.Write("%s\t%s\t", gui.FgHiWhite(key), value)
			}
		}

		ui.MustFlush()
		return nil
	},
}