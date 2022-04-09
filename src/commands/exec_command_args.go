package commands

import (
	"github.com/urfave/cli/v2"
)

type ExecArgs struct {
	Project string
	Command string
}

func ParseAndValidateExecCommandArgs(c *cli.Context) (*ExecArgs, error) {
	project := c.String("project")
	command := c.String("command")

	return &ExecArgs{
		Project: project,
		Command: command,
	}, nil
}
