package main

import (
	"github.com/urfave/cli/v2"
	"strings"
)

var initCommand = &cli.Command{
	Name:      "init",
	Usage:     "Create a .workwork file with default content",
	UsageText: "Running `ww init` will start a wizard that will ask you for URLs for common things like docs, logs, ci, tasks, etc. The more you are able to fill in, the better. You can always add more, update or delete URLs later.",
	Action: func(c *cli.Context) error {
		ui := newUserInterface()

		ui.write("%s", fgHiGreen("Creating a new .workwork file"))

		defaults := map[string]string{
			"contact": "",
			"repo":    "",
			"ci":      "",
			"cd":      "",
			"issues":  "",
			"pulls":   "",
			"tasks":   "",
			"docs":    "",
		}

		envDefaults := map[string]string{
			"logs":       "",
			"monitoring": "",
			"live":       "",
		}

		aborted := false
		for key := range defaults {
			answer := ui.ask("%s '%s' %s", fgHiWhite("Enter a valid URL for"), boldFgHiYellow(key), fgHiWhite("or leave blank to ignore"))

			if answer == "" {
				delete(defaults, key)
				continue
			}

			if isValidUrl(answer) {
				defaults[key] = answer
				continue
			}

			ui.write("'%s' %s", boldFgHiRed(answer), fgHiRed("is not a valid URL"))
			aborted = true
			break // TODO: Try same key again until it's correct instead of exiting
		}

		envs := make([]Environment, 0)
		answer := ui.ask("%s:", fgHiWhite("App environments (space or comma separated e.x.: \"local dev test stage prod\")"))
		trimmedAnswer := strings.TrimSpace(answer)

		var parts []string
		if strings.Contains(trimmedAnswer, ",") {
			parts = strings.Split(trimmedAnswer, ",")
		} else {
			parts = strings.Split(trimmedAnswer, " ")
		}

		for _, part := range parts {
			trimmedPart := strings.TrimSpace(part)
			envs = append(envs, NewEnvironment(trimmedPart, envDefaults))
		}

		for i, env := range envs {
			for key := range env.Urls {
				answer := ui.ask("[%s environment] %s '%s' %s", fgHiMagenta(env.Name), fgHiWhite("Enter a valid URL for"), boldFgHiYellow(key), fgHiWhite("or leave blank to ignore"))

				if answer == "" {
					delete(envs[i].Urls, key)
					continue
				}

				if isValidUrl(answer) {
					envs[i].Urls[key] = answer
					continue
				}

				ui.write("'%s' %s", boldFgHiRed(answer), fgHiRed("is not a valid URL"))
				aborted = true
				break // TODO: Try same key again until it's correct instead of exiting
			}
		}

		if aborted {
			return nil
		}

		ui.write("%s", fgHiGreen("\nRegistered URLs:"))
		for key, value := range defaults {
			ui.write("%s\t%s", fgHiWhite(key), value)
		}

		ui.write("\nYou can use the '%s' command to enter more URLs, or '%s' to read about all commdsn", boldFgHiGreen("set"), boldFgHiGreen("help"))

		file := WorkWorkFile{Urls: defaults, Environments: envs}
		if err := writeWorkWorkFile(&file); err != nil {
			return err
		}

		ui.mustFlush()
		return nil
	},
}
