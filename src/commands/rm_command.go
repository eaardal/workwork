package commands

import (
	"github.com/eaardal/workwork/src/gui"
	"github.com/eaardal/workwork/src/ww"
	"github.com/urfave/cli/v2"
)

var RMCommand = &cli.Command{
	Name:      "rm",
	Usage:     "Remove the URL for the given key",
	UsageText: "Remove a single URL: `ww rm {key}`. Remove many at once: `ww rm {key1} {key2} {key3}`",
	Action: func(c *cli.Context) error {
		ui := gui.NewUserInterface()
		defer ui.MustFlush()

		args, err := ParseAndValidateRMCommandArgs(c)
		if err != nil {
			return err
		}

		wwFile, err := ww.ReadWorkWorkYaml()
		if err != nil {
			return err
		}

		for _, keyToDelete := range args.UrlKeysToBeDeleted {
			itemExists := false

			for existingKey, existingUrl := range wwFile.GlobalUrls {
				if existingKey == keyToDelete {
					delete(wwFile.GlobalUrls, existingKey)
					ui.Write("%s '%s' (%s)", gui.FgHiGreen("Removed"), gui.BoldFgHiYellow(keyToDelete), existingUrl)
					itemExists = true
					break
				}
			}

			if !itemExists {
				ui.Write("%s '%s'", gui.FgHiRed("Found no URL with key"), gui.BoldFgHiYellow(keyToDelete))
			}
		}

		if err := ww.WriteWorkWorkYaml(wwFile); err != nil {
			return err
		}

		return nil
	},
}
