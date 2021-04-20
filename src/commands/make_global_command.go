package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"path"
	"workwork/src/gui"
	"workwork/src/utils"
	"workwork/src/ww"
)

var MakeGlobalCommand = &cli.Command{
	Name:      "make-global",
	Usage:     "Store a .workwork.yaml globally instead of inside a repository. Only moves existing .workwork.yaml files.",
	UsageText: "Example: `ww move-global {path-to-workworkyaml}`. Only works on existing .workwork.yaml files. See `ww init --global` to make a new .workwork.yaml file and store it globally",
	Action: func(c *cli.Context) error {
		ui := gui.NewUserInterface()

		args, err := ParseAndValidateMakeGlobalCommandArgs(c)
		if err != nil {
			return err
		}

		globalRoot, err := ww.EnsureGlobalDirectoryExists()
		if err != nil {
			return err
		}

		globalRepoPath := path.Join(globalRoot, args.RepositoryName)
		if err := utils.CreateDirectoryIfNotExists(globalRepoPath); err != nil {
			return err
		}

		repositoryWorkWorkYamlPath := path.Join(args.RepositoryPath, ".workwork.yaml")
		if !utils.FileExists(repositoryWorkWorkYamlPath) {
			return fmt.Errorf("did not find an existing .workwork.yaml file at %s. You might want to run `ww init --global` to create a new .workwork.yaml file globally for this repository, or check that the provided path is correct", repositoryWorkWorkYamlPath)
		}

		globalWorkWorkYamlPath := path.Join(globalRepoPath, ".workwork.yaml")

		if err := os.Rename(repositoryWorkWorkYamlPath, globalWorkWorkYamlPath); err != nil {
			return err
		}

		ui.Write("%s", gui.FgHiGreen(fmt.Sprintf("Moved '%s' to '%s'", repositoryWorkWorkYamlPath, globalWorkWorkYamlPath)))

		ui.MustFlush()
		return nil
	},
}
