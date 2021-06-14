package commands

import (
	"fmt"
	"github.com/eaardal/workwork/src/utils"
	"github.com/eaardal/workwork/src/ww"
	"github.com/urfave/cli/v2"
	"strings"
)

type GetArgs struct {
	UrlKeys          map[string][]string
	WorkingDirectory string
	Global           bool
}

func ParseAndValidateGetCommandArgs(c *cli.Context) (*GetArgs, error) {
	urls := make(map[string][]string, 0)
	args := c.Args().Slice()

	for _, arg := range args {
		trimmed := strings.TrimSpace(arg)
		if strings.Contains(trimmed, ".") {
			parts := strings.Split(trimmed, ".")
			if len(parts) == 2 {
				environmentName := parts[0]
				urlKey := parts[1]
				if urls[environmentName] == nil {
					urls[environmentName] = make([]string, 0)
				}
				urls[environmentName] = append(urls[environmentName], urlKey)
			} else {
				return nil, fmt.Errorf("unexpected format for key '%s'", trimmed)
			}
		} else {
			if urls[ww.Global] == nil {
				urls[ww.Global] = make([]string, 0)
			}
			urls[ww.Global] = append(urls[ww.Global], trimmed)
		}
	}

	useGlobalWorkWorkYamlFile := c.Bool(utils.GlobalFlagName)

	wd, err := utils.ResolveWorkingDirectory(c, useGlobalWorkWorkYamlFile)
	if err != nil {
		return nil, fmt.Errorf("unable to resolve wd: %v", err)
	}

	return &GetArgs{
		UrlKeys:          urls,
		WorkingDirectory: wd,
		Global:           useGlobalWorkWorkYamlFile,
	}, nil
}
