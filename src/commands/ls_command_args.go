package commands

import (
	"fmt"
	"github.com/eaardal/workwork/src/utils"
	"github.com/urfave/cli/v2"
)

type LSCommandArgs struct {
	WorkingDirectory string
	Global           bool
}

func ParseAndValidateLSCommandArgs(c *cli.Context) (*LSCommandArgs, error) {
	useGlobalWorkWorkYamlFile := c.Bool(utils.GlobalFlagName)

	wd, err := utils.ResolveWorkingDirectory(c, useGlobalWorkWorkYamlFile)
	if err != nil {
		return nil, fmt.Errorf("unable to resolve wd: %v", err)
	}

	return &LSCommandArgs{
		WorkingDirectory: wd,
		Global:           useGlobalWorkWorkYamlFile,
	}, nil
}
