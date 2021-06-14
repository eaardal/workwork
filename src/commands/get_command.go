package commands

import (
	"github.com/eaardal/workwork/src/gui"
	"github.com/eaardal/workwork/src/utils"
	"github.com/eaardal/workwork/src/ww"
	"github.com/urfave/cli/v2"
)

var GetCommand = &cli.Command{
	Name:      "get",
	Usage:     "Show the URL for a specific key",
	UsageText: "Example: `ww get docs`, `ww get docs prod.logs dev.logs`",
	Flags: []cli.Flag{
		utils.BuildWorkingDirectoryFlag(),
		utils.BuildGlobalFlag(),
	},
	Action: func(c *cli.Context) error {
		ui := gui.NewUserInterface()
		defer ui.MustFlush()

		args, err := ParseAndValidateGetCommandArgs(c)
		if err != nil {
			return err
		}

		wwYaml, err := ww.ReadWorkWorkYaml(args.WorkingDirectory)
		if err != nil {
			return err
		}

		for environmentName, environmentUrlKeys := range args.UrlKeys {
			wwFileUrls, err := wwYaml.GetUrls(environmentName)
			if err != nil {
				return err
			}

			ui.Write("%s", gui.FgHiGreen(environmentName))

			for _, urlKey := range environmentUrlKeys {
				itemExists := false

				for fileUrlKey, url := range wwFileUrls {
					if urlKey == fileUrlKey {
						ui.Write("%s\t%s\t", fileUrlKey, url)
						itemExists = true
						break
					}
				}

				if !itemExists {
					ui.Write("%s\t%s\t", gui.BoldFgHiRed(urlKey), gui.FgHiRed("not found"))
				}
			}
		}

		if err := ww.WriteWorkWorkYaml(args.WorkingDirectory, wwYaml); err != nil {
			return err
		}

		return nil
	},
}
