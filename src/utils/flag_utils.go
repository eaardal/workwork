package utils

import (
	"github.com/urfave/cli/v2"
	"os"
)

const (
	WorkingDirectoryFlagName = "working-directory"
	GlobalFlagName           = "global"
)

func WorkingDirectoryFlag() cli.Flag {
	return &cli.StringFlag{
		Name:        WorkingDirectoryFlagName,
		Aliases:     []string{"wd"},
		Usage:       "The full path to the working directory for the command. If not set, the current working directory is used.",
		EnvVars:     []string{"WORKWORK_WD"},
		DefaultText: ".",
	}
}

func ResolveWorkingDirectory(c *cli.Context, useGlobalWorkWorkYamlFile bool) (string, error) {
	if useGlobalWorkWorkYamlFile {
		home := os.Getenv("HOME")
		return home, nil
	}

	workingDirectory := c.String(WorkingDirectoryFlagName)
	if workingDirectory == "" {
		return os.Getwd()
	}

	return workingDirectory, nil
}

func GlobalFlag() cli.Flag {
	return &cli.BoolFlag{
		Name:    GlobalFlagName,
		Aliases: []string{"g"},
		Usage:   "Use the global .workwork.yaml for your machine user profile file instead of a local/project/repository one",
	}
}
