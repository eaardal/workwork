package commands

import (
	"fmt"
	"github.com/eaardal/workwork/src/utils"
	"github.com/eaardal/workwork/src/validation"
	"github.com/urfave/cli/v2"
	"strings"
)

type GoToArgs struct {
	UrlKey           string
	Environment      string
	HasEnvironment   bool
	WorkingDirectory string
}

func ParseAndValidateGoToCommandArgs(c *cli.Context) (*GoToArgs, error) {
	args := c.Args()

	var keyArg string
	var envArg string

	if args.Len() == 2 {
		envArg = args.Get(0)
		keyArg = args.Get(1)
	} else {
		keyArg = c.Args().Get(0)
		if strings.Contains(keyArg, ".") {
			parts := strings.Split(keyArg, ".")
			if len(parts) == 2 {
				envArg = parts[0]
				keyArg = parts[1]
			} else {
				return nil, fmt.Errorf("unexpected format for key '%s'", keyArg)
			}
		} else {
			envArg = ""
		}
	}

	if !validation.IsValidKey(keyArg) {
		return nil, fmt.Errorf("invalid key '%s'", keyArg)
	}

	wd, err := utils.ResolveWorkingDirectory(c)
	if err != nil {
		return nil, fmt.Errorf("unable to resolve wd: %v", err)
	}

	return &GoToArgs{
		UrlKey:           keyArg,
		Environment:      envArg,
		HasEnvironment:   envArg != "",
		WorkingDirectory: wd,
	}, nil
}
