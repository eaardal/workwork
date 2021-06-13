package commands

import (
	"github.com/eaardal/workwork/src/gui"
	"github.com/eaardal/workwork/src/utils"
	"github.com/eaardal/workwork/src/ww"
	"github.com/urfave/cli/v2"
)

var LSCommand = &cli.Command{
	Name:      "ls",
	Usage:     "List all the registered URLs",
	UsageText: "Example: `ww ls`",
	Flags: []cli.Flag{
		utils.BuildWorkingDirectoryFlag(),
	},
	Action: func(c *cli.Context) error {
		args, err := ParseAndValidateLSCommandArgs(c)
		if err != nil {
			return err
		}

		wwYaml, err := ww.ReadWorkWorkYaml(args.WorkingDirectory)
		if err != nil {
			return err
		}

		ui := gui.NewUserInterface()
		defer ui.MustFlush()

		ui.Write("%s", gui.FgHiGreen("global"))

		for key, value := range wwYaml.GlobalUrls {
			ui.Write("%s\t%s\t", key, value)
		}

		for _, env := range wwYaml.Environments {
			ui.Write("\n%s", gui.FgHiGreen(env.Name))

			for key, value := range env.EnvironmentUrls {
				ui.Write("%s\t%s\t", key, value)
			}
		}

		return nil
	},
}
