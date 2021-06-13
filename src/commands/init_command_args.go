package commands

import (
	"fmt"
	"github.com/eaardal/workwork/src/utils"
	"github.com/urfave/cli/v2"
)

type InitCommandArgs struct {
	WorkingDirectory string
}

func ParseAndValidateInitCommandArgs(c *cli.Context) (*InitCommandArgs, error) {
	wd, err := utils.ResolveWorkingDirectory(c)
	if err != nil {
		return nil, fmt.Errorf("unable to resolve wd: %v", err)
	}

	return &InitCommandArgs{WorkingDirectory: wd}, nil
}
