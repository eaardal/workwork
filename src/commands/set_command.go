package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
	"workwork/src/gui"
	"workwork/src/validation"
	"workwork/src/ww"
)

var SetCommand = &cli.Command{
	Name:      "set",
	Usage:     "Set a new URL for an existing key, or add a new URL if the key doesn't exist",
	UsageText: "Add or update single item: `ww set key=url`. Add or update many at once: `ww set key=url key=url key=url`. If the key exists, the URL will be updated. If the key doesn't exist, the URL will be added. Keys must consist of a-z lower cased letters only and use snake_case if you need spacing.",
	Action: func(c *cli.Context) error {
		printer := gui.NewUserInterface()

		wwFile, err := ww.ReadWorkWorkFile()
		if err != nil {
			return err
		}

		newUrls := make(map[string]string, 0)

		for _, item := range c.Args().Slice() {
			if !strings.Contains(item, "=") {
				printer.Write("'%s' %s", gui.BoldFgHiRed(item), gui.FgHiRed("is not a valid key=value pair"))
				return nil
			}

			parts := strings.Split(item, "=")
			if len(parts) != 2 {
				printer.Write("'%s' %s %s %s", gui.BoldFgHiRed(item), gui.FgHiRed("was split into"), gui.BoldFgHiRed(fmt.Sprintf("%d", len(parts))), gui.FgHiRed("parts, but expected to find 2 parts"))
				return nil
			}

			key := parts[0]
			if !validation.IsValidKey(key) {
				printer.Write("'%s' %s", gui.BoldFgHiRed(key), gui.FgHiRed("is not a valid key. Must be lowercase and snake_case"))
				return nil
			}

			value := parts[1]

			if !validation.IsValidUrl(value) {
				printer.Write("'%s' %s", gui.BoldFgHiRed(value), gui.FgHiRed("is not a valid URL"))
				return nil
			}

			newUrls[key] = value
		}

		for newKey, newValue := range newUrls {
			itemExists := false

			for existingKey, existingValue := range wwFile.Urls {
				if newKey == existingKey && existingValue != newValue {
					wwFile.Urls[existingKey] = newValue
					printer.Write("Updated '%s' to '%s' (was '%s')", gui.BoldFgHiYellow(existingKey), gui.FgHiWhite(newValue), existingValue)
					itemExists = true
					break
				}
			}

			if !itemExists {
				wwFile.Urls[newKey] = newValue
				printer.Write("Added '%s' with URL '%s'", gui.BoldFgHiYellow(newKey), gui.FgHiWhite(newValue))
			}
		}

		if err := ww.WriteWorkWorkFile(wwFile); err != nil {
			return err
		}

		return nil
	},
}
