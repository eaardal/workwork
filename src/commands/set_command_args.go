package commands

import (
	"fmt"
	"github.com/eaardal/workwork/src/utils"
	"github.com/eaardal/workwork/src/validation"
	"github.com/eaardal/workwork/src/ww"
	"github.com/urfave/cli/v2"
	"strings"
)

type SetArgs struct {
	Urls             ww.Urls
	WorkingDirectory string
	Global           bool
}

func ParseAndValidateSetCommandArgs(c *cli.Context) (*SetArgs, error) {
	urls := make(ww.Urls, 0)

	for _, item := range c.Args().Slice() {
		if !strings.Contains(item, "=") {
			return nil, fmt.Errorf("'%s' is not a valid key=value pair", item)
		}

		parts := strings.Split(item, "=")
		if len(parts) != 2 {
			return nil, fmt.Errorf("'%s' was split into %d parts, but expected to find 2 parts", item, len(parts))
		}

		key := parts[0]
		if !validation.IsValidKey(key) {
			return nil, fmt.Errorf("'%s' is not a valid key. Must be lowercase and snake_case", key)
		}

		value := parts[1]

		if !validation.IsValidUrl(value) {
			return nil, fmt.Errorf("'%s' is not a valid URL", value)
		}

		urls[key] = value
	}

	useGlobalWorkWorkYamlFile := c.Bool(utils.GlobalFlag)

	wd, err := utils.ResolveWorkingDirectory(c, useGlobalWorkWorkYamlFile)
	if err != nil {
		return nil, fmt.Errorf("unable to resolve wd: %v", err)
	}

	return &SetArgs{
		Urls:             urls,
		WorkingDirectory: wd,
		Global:           useGlobalWorkWorkYamlFile,
	}, nil
}
