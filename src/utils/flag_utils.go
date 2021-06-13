package utils

import (
	"github.com/urfave/cli/v2"
	"os"
)

func BuildWorkingDirectoryFlag() cli.Flag {
	return &cli.StringFlag{
		Name:        "working-directory",
		Aliases:     []string{"wd"},
		Usage:       "The full path to the working directory for the command. If not set, the current working directory is used.",
		EnvVars:     []string{"WORKWORK_WD"},
		DefaultText: ".",
	}
}

func ResolveWorkingDirectory(c *cli.Context) (string, error) {
	workingDirectory := c.String("working-directory")
	if workingDirectory == "" {
		return os.Getwd()
	}
	return workingDirectory, nil
}
