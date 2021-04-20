package commands

import (
	"github.com/urfave/cli/v2"
	"workwork/src/gui"
	"workwork/src/ww"
)

var GetCommand = &cli.Command{
	Name:      "get",
	Usage:     "Show the URL for a specific key",
	UsageText: "Example: `ww get docs`, `ww get docs prod.logs dev.logs`",
	Action: func(c *cli.Context) error {
		ui := gui.NewUserInterface()

		wwFile, err := ww.ReadWorkWorkFile()
		if err != nil {
			return err
		}

		args, err := ParseAndValidateGetCommandArgs(c)
		if err != nil {
			return err
		}

		for environmentName, environmentUrlKeys := range args.UrlKeys {
			wwFileUrls, err := wwFile.GetUrls(environmentName)
			if err != nil {
				return err
			}

			ui.Write("%s", gui.FgHiGreen(environmentName))

			for _, urlKey := range environmentUrlKeys {
				itemExists := false

				for fileUrlKey, url := range wwFileUrls {
					if urlKey == fileUrlKey {
						ui.Write("%s\t%s\t", gui.FgHiWhite(fileUrlKey), url)
						itemExists = true
						break
					}
				}

				if !itemExists {
					ui.Write("%s\t%s\t", gui.BoldFgHiRed(urlKey), gui.FgHiRed("not found"))
				}
			}
		}

		if err := ww.WriteWorkWorkFile(wwFile); err != nil {
			return err
		}

		ui.MustFlush()
		return nil
	},
}
