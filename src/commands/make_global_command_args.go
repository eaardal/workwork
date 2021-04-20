package commands

import "github.com/urfave/cli/v2"

type MakeGlobalArgs struct {
	RepositoryPath string
	RepositoryName string
}

func ParseAndValidateMakeGlobalCommandArgs(c *cli.Context) (*MakeGlobalArgs, error) {
	return &MakeGlobalArgs{}, nil
}
