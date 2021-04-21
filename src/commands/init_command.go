package commands

import (
	"github.com/eaardal/workwork/src/gui"
	"github.com/eaardal/workwork/src/validation"
	"github.com/eaardal/workwork/src/ww"
	"github.com/urfave/cli/v2"
	"strings"
)

var InitCommand = &cli.Command{
	Name:      "init",
	Usage:     "Create a .workwork file with default content",
	UsageText: "Running `ww init` will start a wizard that will ask you for URLs for common things like docs, logs, ci, tasks, etc. The more you are able to fill in, the better. You can always add more, update or delete URLs later.",
	Action: func(c *cli.Context) error {
		ui := gui.NewUserInterface()

		ui.Write("%s", gui.FgHiGreen("Creating a new .workwork file"))

		globalUrls := map[string]string{
			"contact": "",
			"repo":    "",
			"ci":      "",
			"cd":      "",
			"issues":  "",
			"pulls":   "",
			"tasks":   "",
			"docs":    "",
		}

		environmentUrls := map[string]string{
			"logs":       "",
			"monitoring": "",
			"live":       "",
		}

		aborted := fillGlobalUrls(ui, globalUrls)
		if aborted {
			return nil
		}

		envs, aborted := fillEnvironmentUrls(ui, environmentUrls)
		if aborted {
			return nil
		}

		printUrls(ui, globalUrls, envs)

		file := ww.WorkWorkYaml{GlobalUrls: globalUrls, Environments: envs}
		if err := ww.WriteWorkWorkYaml(&file); err != nil {
			return err
		}

		ui.MustFlush()
		return nil
	},
}

func fillGlobalUrls(ui gui.UserInterface, urls map[string]string) (aborted bool) {
	for key := range urls {
		answer := ui.Ask("%s '%s' %s", gui.FgHiWhite("Enter a valid URL for"), gui.BoldFgHiYellow(key), gui.FgHiWhite("or leave blank to ignore"))

		if answer == "" {
			delete(urls, key)
			continue
		}

		if validation.IsValidUrl(answer) {
			urls[key] = answer
			continue
		}

		ui.Write("'%s' %s", gui.BoldFgHiRed(answer), gui.FgHiRed("is not a valid URL"))
		aborted = true
		break // TODO: Try same key again until it's correct instead of exiting
	}

	return aborted
}

func fillEnvironmentUrls(ui gui.UserInterface, urls map[string]string) (envs []ww.Environment, aborted bool) {
	answer := ui.Ask("%s:", gui.FgHiWhite("App environments (space or comma separated e.x.: \"local dev test stage prod\")"))
	trimmedAnswer := strings.TrimSpace(answer)

	var parts []string
	if strings.Contains(trimmedAnswer, ",") {
		parts = strings.Split(trimmedAnswer, ",")
	} else {
		parts = strings.Split(trimmedAnswer, " ")
	}

	for _, part := range parts {
		trimmedPart := strings.TrimSpace(part)
		lower := strings.ToLower(trimmedPart)
		// TODO: Validate environment name
		envs = append(envs, ww.NewEnvironment(lower, urls))
	}

	for i, env := range envs {
		for key := range env.EnvironmentUrls {
			answer := ui.Ask("[%s environment] %s '%s' %s", gui.FgHiMagenta(env.Name), gui.FgHiWhite("Enter a valid URL for"), gui.BoldFgHiYellow(key), gui.FgHiWhite("or leave blank to ignore"))

			if answer == "" {
				delete(envs[i].EnvironmentUrls, key)
				continue
			}

			if validation.IsValidUrl(answer) {
				envs[i].EnvironmentUrls[key] = answer
				continue
			}

			ui.Write("'%s' %s", gui.BoldFgHiRed(answer), gui.FgHiRed("is not a valid URL"))
			aborted = true
			break // TODO: Try same key again until it's correct instead of exiting
		}
	}

	return envs, aborted
}

func printUrls(ui gui.UserInterface, globalUrls ww.Urls, environmentUrls []ww.Environment) {
	ui.Write("%s", gui.FgHiGreen("\nglobal"))

	for key, value := range globalUrls {
		ui.Write("%s\t%s\t", gui.FgHiWhite(key), value)
	}

	for _, env := range environmentUrls {
		ui.Write("\n%s", gui.FgHiGreen(env.Name))

		for key, value := range env.EnvironmentUrls {
			ui.Write("%s\t%s\t", gui.FgHiWhite(key), value)
		}
	}

	ui.Write("\nYou can use the '%s' command to enter more URLs, or '%s' to read about all commands", gui.BoldFgHiGreen("set"), gui.BoldFgHiGreen("help"))
}
