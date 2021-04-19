package main

import (
	"github.com/urfave/cli/v2"
)

var rm = &cli.Command{
	Name:  "rm",
	Usage: "Remove a URL",
	Action: func(c *cli.Context) error {
		printer := newUserInterface()

		ww, err := readWorkWorkFile()
		if err != nil {
			return err
		}

		toBeDeleted := make([]string, 0)

		for _, item := range c.Args().Slice() {
			toBeDeleted = append(toBeDeleted, item)
		}

		for _, keyToDelete := range toBeDeleted {
			itemExists := false

			for key, url := range ww.Urls {
				if key == keyToDelete {
					delete(ww.Urls, key)
					printer.write("%s '%s' (%s)", hiGreen("Removed"), boldHiYellow(keyToDelete), hiWhite(url))
					itemExists = true
					break
				}
			}

			if !itemExists {
				printer.write("%s '%s'", hiRed("Found no URL with key"), boldHiYellow(keyToDelete))
			}
		}

		if err := writeWorkWorkFile(ww); err != nil {
			return err
		}

		return nil
	},
}
