package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

type RMArgs struct {
	UrlKeysToBeDeleted []string
}

func ParseAndValidateRMCommandArgs(c *cli.Context) (*RMArgs, error) {
	if c.NArg() == 0 {
		return nil, fmt.Errorf("no url keys specified. See `ww rm --help` for correct usage")
	}

	return &RMArgs{
		UrlKeysToBeDeleted: c.Args().Slice(),
	}, nil
}
