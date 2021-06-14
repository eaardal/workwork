package commands

import (
	"fmt"
	"github.com/eaardal/workwork/src/utils"
	"github.com/urfave/cli/v2"
)

type InitCommandArgs struct {
	WorkingDirectory string
	Global           bool
}

func ParseAndValidateInitCommandArgs(c *cli.Context) (*InitCommandArgs, error) {
	useGlobalWorkWorkYamlFile := c.Bool(utils.GlobalFlag)

	wd, err := utils.ResolveWorkingDirectory(c, useGlobalWorkWorkYamlFile)
	if err != nil {
		return nil, fmt.Errorf("unable to resolve wd: %v", err)
	}

	return &InitCommandArgs{
		WorkingDirectory: wd,
		Global:           useGlobalWorkWorkYamlFile,
	}, nil
}
