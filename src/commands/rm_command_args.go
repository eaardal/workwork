package commands

import (
	"fmt"
	"github.com/eaardal/workwork/src/utils"
	"github.com/urfave/cli/v2"
)

type RMArgs struct {
	UrlKeysToBeDeleted []string
	WorkingDirectory   string
}

func ParseAndValidateRMCommandArgs(c *cli.Context) (*RMArgs, error) {
	if c.NArg() == 0 {
		return nil, fmt.Errorf("no url keys specified. See `ww rm --help` for correct usage")
	}

	wd, err := utils.ResolveWorkingDirectory(c)
	if err != nil {
		return nil, fmt.Errorf("unable to resolve wd: %v", err)
	}

	return &RMArgs{
		UrlKeysToBeDeleted: c.Args().Slice(),
		WorkingDirectory:   wd,
	}, nil
}
