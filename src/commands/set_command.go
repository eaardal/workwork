package commands

import (
	"github.com/eaardal/workwork/src/gui"
	"github.com/eaardal/workwork/src/utils"
	"github.com/eaardal/workwork/src/ww"
	"github.com/urfave/cli/v2"
)

var SetCommand = &cli.Command{
	Name:      "set",
	Usage:     "Set a new URL for an existing key, or add a new URL if the key doesn't exist",
	UsageText: "Add or update single item: `ww set key=url`. Add or update many at once: `ww set key=url key=url key=url`. If the key exists, the URL will be updated. If the key doesn't exist, the URL will be added. Keys must consist of a-z lower cased letters only and use snake_case if you need spacing.",
	Flags: []cli.Flag{
		utils.WorkingDirectoryFlag(),
		utils.GlobalFlag(),
	},
	Action: func(c *cli.Context) error {
		ui := gui.NewUserInterface()
		defer ui.MustFlush()

		args, err := ParseAndValidateSetCommandArgs(c)
		if err != nil {
			return err
		}

		wwYaml, err := ww.ReadWorkWorkYaml(args.WorkingDirectory)
		if err != nil {
			return err
		}

		for urlKey, url := range args.Urls {
			itemExists := false

			for existingKey, existingValue := range wwYaml.GlobalUrls {
				if urlKey == existingKey && existingValue != url {
					wwYaml.GlobalUrls[existingKey] = url
					ui.Write("Updated '%s' to '%s' (was '%s')", gui.BoldFgHiYellow(existingKey), url, existingValue)
					itemExists = true
					break
				}
			}

			if !itemExists {
				wwYaml.GlobalUrls[urlKey] = url
				ui.Write("Added '%s' with URL '%s'", gui.BoldFgHiYellow(urlKey), url)
			}
		}

		if err := ww.WriteWorkWorkYaml(args.WorkingDirectory, wwYaml); err != nil {
			return err
		}

		return nil
	},
}
