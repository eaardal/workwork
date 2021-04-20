package commands

import (
	"github.com/urfave/cli/v2"
	"workwork/src/gui"
	"workwork/src/ww"
)

var RMCommand = &cli.Command{
	Name:      "rm",
	Usage:     "Remove the URL for the given key",
	UsageText: "Remove a single URL: `ww rm {key}`. Remove many at once: `ww rm {key1} {key2} {key3}`",
	Action: func(c *cli.Context) error {
		printer := gui.NewUserInterface()

		wwFile, err := ww.ReadWorkWorkFile()
		if err != nil {
			return err
		}

		toBeDeleted := make([]string, 0)

		for _, item := range c.Args().Slice() {
			toBeDeleted = append(toBeDeleted, item)
		}

		for _, keyToDelete := range toBeDeleted {
			itemExists := false

			for key, url := range wwFile.Urls {
				if key == keyToDelete {
					delete(wwFile.Urls, key)
					printer.Write("%s '%s' (%s)", gui.FgHiGreen("Removed"), gui.BoldFgHiYellow(keyToDelete), gui.FgHiWhite(url))
					itemExists = true
					break
				}
			}

			if !itemExists {
				printer.Write("%s '%s'", gui.FgHiRed("Found no URL with key"), gui.BoldFgHiYellow(keyToDelete))
			}
		}

		if err := ww.WriteWorkWorkFile(wwFile); err != nil {
			return err
		}

		return nil
	},
}
