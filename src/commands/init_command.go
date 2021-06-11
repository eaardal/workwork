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
		defer ui.MustFlush()

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

		fillGlobalUrls(ui, globalUrls)
		envs := fillEnvironmentUrls(ui, environmentUrls)

		printUrls(ui, globalUrls, envs)

		file := ww.WorkWorkYaml{GlobalUrls: globalUrls, Environments: envs}
		if err := ww.WriteWorkWorkYaml(&file); err != nil {
			return err
		}

		return nil
	},
}

func fillGlobalUrls(ui gui.UserInterface, urls map[string]string) {
	for key := range urls {
		for {
			answer := ui.Ask("Enter a valid URL for '%s' or leave blank to ignore", gui.BoldFgHiYellow(key))

			if answer == "" {
				delete(urls, key)
				break
			}

			if validation.IsValidUrl(answer) {
				urls[key] = answer
				break
			}

			ui.Write("'%s' %s", gui.BoldFgHiRed(answer), gui.FgHiRed("is not a valid URL"))
		}
	}
}

func fillEnvironmentUrls(ui gui.UserInterface, urls map[string]string) (envs []ww.Environment) {
	answer := ui.Ask("App environments (space or comma separated e.x.: \"local dev test stage prod\"):")
	trimmedAnswer := strings.TrimSpace(answer)

	if trimmedAnswer == "" {
		ui.Write("No environments added. You can always add environments manually in .workwork.yaml later.")
		return envs
	}

	var parts []string
	if strings.Contains(trimmedAnswer, ",") {
		parts = strings.Split(trimmedAnswer, ",")
	} else {
		parts = strings.Split(trimmedAnswer, " ")
	}

	for _, part := range parts {
		trimmedPart := strings.TrimSpace(part)
		lower := strings.ToLower(trimmedPart)
		envs = append(envs, ww.NewEnvironment(lower, urls))
	}

	for i, env := range envs {
		for key := range env.EnvironmentUrls {
			for {
				answer := ui.Ask("[%s environment] Enter a valid URL for '%s' or leave blank to ignore", gui.FgHiMagenta(env.Name), gui.BoldFgHiYellow(key))

				if answer == "" {
					delete(envs[i].EnvironmentUrls, key)
					break
				}

				if validation.IsValidUrl(answer) {
					envs[i].EnvironmentUrls[key] = answer
					break
				}

				ui.Write("'%s' %s", gui.BoldFgHiRed(answer), gui.FgHiRed("is not a valid URL"))
			}
		}
	}

	return envs
}

func printUrls(ui gui.UserInterface, globalUrls ww.Urls, environmentUrls []ww.Environment) {
	ui.Write("%s", gui.FgHiGreen("\nglobal"))

	for key, value := range globalUrls {
		ui.Write("%s\t%s\t", key, value)
	}

	for _, env := range environmentUrls {
		ui.Write("\n%s", gui.FgHiGreen(env.Name))

		for key, value := range env.EnvironmentUrls {
			ui.Write("%s\t%s\t", key, value)
		}
	}

	ui.Write("\nYou can use the '%s' command to enter more URLs, or '%s' to read about all commands", gui.BoldFgHiGreen("set"), gui.BoldFgHiGreen("help"))
}
