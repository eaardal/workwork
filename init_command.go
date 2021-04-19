package main

import (
	"github.com/urfave/cli/v2"
)

var initCommand = &cli.Command{
	Name:      "init",
	Usage:     "Create a .workwork file with default content",
	UsageText: "Running `ww init` will start a wizard that will ask you for URLs for common things like docs, logs, ci, tasks, etc. The more you are able to fill in, the better. You can always add more, update or delete URLs later.",
	Action: func(c *cli.Context) error {
		ui := newUserInterface()

		ui.write("%s", hiGreen("Creating a new .workwork file"))

		defaults := map[string]string{
			"contact":    "",
			"repo":       "",
			"ci":         "",
			"cd":         "",
			"issues":     "",
			"pulls":      "",
			"tasks":      "",
			"docs":       "",
			"logs":       "",
			"monitoring": "",
		}

		aborted := false
		for key := range defaults {
			answer := ui.askUser("%s '%s' %s", hiWhite("Enter a valid URL for"), boldHiYellow(key), hiWhite("or leave blank to ignore"))

			if answer == "" {
				delete(defaults, key)
				continue
			}

			if isValidUrl(answer) {
				defaults[key] = answer
				continue
			}

			ui.write("'%s' %s", boldHiRed(answer), hiRed("is not a valid URL"))
			aborted = true
			break // TODO: Try same key again until it's correct instead of exiting
		}

		if aborted {
			return nil
		}

		ui.write("%s", hiGreen("\nRegistered URLs:"))
		for key, value := range defaults {
			ui.write("%s\t%s", hiWhite(key), value)
		}

		ui.write("\nYou can use the '%s' command to enter more URLs, or '%s' to read about all commdsn", boldHiGreen("set"), boldHiGreen("help"))

		file := WorkWorkFile{Urls: defaults}
		if err := writeWorkWorkFile(&file); err != nil {
			return err
		}

		ui.mustFlush()
		return nil
	},
}
