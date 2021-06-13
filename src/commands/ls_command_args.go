package commands

import (
	"fmt"
	"github.com/eaardal/workwork/src/utils"
	"github.com/urfave/cli/v2"
)

type LSCommandArgs struct {
	WorkingDirectory string
}

func ParseAndValidateLSCommandArgs(c *cli.Context) (*LSCommandArgs, error) {
	wd, err := utils.ResolveWorkingDirectory(c)
	if err != nil {
		return nil, fmt.Errorf("unable to resolve wd: %v", err)
	}

	return &LSCommandArgs{WorkingDirectory: wd}, nil
}
