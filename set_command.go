package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
)

var setCommand = &cli.Command{
	Name:      "set",
	Usage:     "Set a new URL for an existing key, or add a new URL if the key doesn't exist",
	UsageText: "Add or update single item: `ww set key=url`. Add or update many at once: `ww set key=url key=url key=url`. If the key exists, the URL will be updated. If the key doesn't exist, the URL will be added. Keys must consist of a-z lower cased letters only and use snake_case if you need spacing.",
	Action: func(c *cli.Context) error {
		printer := newUserInterface()

		ww, err := readWorkWorkFile()
		if err != nil {
			return err
		}

		newUrls := make(map[string]string, 0)

		for _, item := range c.Args().Slice() {
			if !strings.Contains(item, "=") {
				printer.write("'%s' %s", boldHiRed(item), hiRed("is not a valid key=value pair"))
				return nil
			}

			parts := strings.Split(item, "=")
			if len(parts) != 2 {
				printer.write("'%s' %s %s %s", boldHiRed(item), hiRed("was split into"), boldHiRed(fmt.Sprintf("%d", len(parts))), hiRed("parts, but expected to find 2 parts"))
				return nil
			}

			key := parts[0]
			if !isValidKey(key) {
				printer.write("'%s' %s", boldHiRed(key), hiRed("is not a valid key. Must be lowercase and snake_case"))
				return nil
			}

			value := parts[1]

			if !isValidUrl(value) {
				printer.write("'%s' %s", boldHiRed(value), hiRed("is not a valid URL"))
				return nil
			}

			newUrls[key] = value
		}

		for newKey, newValue := range newUrls {
			itemExists := false

			for existingKey, existingValue := range ww.Urls {
				if newKey == existingKey && existingValue != newValue {
					ww.Urls[existingKey] = newValue
					printer.write("Updated '%s' to '%s' (was '%s')", boldHiYellow(existingKey), hiWhite(newValue), existingValue)
					itemExists = true
					break
				}
			}

			if !itemExists {
				ww.Urls[newKey] = newValue
				printer.write("Added '%s' with URL '%s'", boldHiYellow(newKey), hiWhite(newValue))
			}
		}

		if err := writeWorkWorkFile(ww); err != nil {
			return err
		}

		return nil
	},
}
