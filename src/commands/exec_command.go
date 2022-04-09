package commands

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/eaardal/workwork/src/gui"
	"github.com/eaardal/workwork/src/ww"
	"github.com/urfave/cli/v2"
)

var ExecCommand = &cli.Command{
	Name:      "exec",
	Usage:     "Execute a CLI command against all repos in a project",
	UsageText: "Example: `ww exec --project {PROJECT_ID} --command {SHELL COMMAND}`",
	Aliases:   []string{"e"},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "project",
			Aliases: []string{"p"},
			Usage:   "The project containing the repositories to execute the command against",
			EnvVars: []string{"WORKWORK_PROJECT"},
		},
		&cli.StringFlag{
			Name:    "command",
			Aliases: []string{"c"},
			Usage:   "The CLI command to execute",
		},
		&cli.BoolFlag{
			Name: "make",
		},
	},
	Action: func(c *cli.Context) error {
		ui := gui.NewUserInterface()
		defer ui.MustFlush()

		args, err := ParseAndValidateExecCommandArgs(c)
		if err != nil {
			return err
		}

		makeData := c.Bool("make")
		if makeData {
			err := ww.WriteGlobalWorkWorkYaml(
				&ww.GlobalWorkWorkYaml{
					Projects: []ww.Project{
						{
							Name: "reklame",
							Repos: []ww.Repository{
								{
									Name: "rp-ad-api-gateway",
									Path: "~/dev/git/tv2/rp-ad-api-gateway",
								},
								{
									Name: "rp-segment-api",
									Path: "~/dev/git/tv2/rp-segment-api",
								},
							},
						},
					},
				},
			)
			if err != nil {
				return fmt.Errorf("failed to write test data: %v", err)
			}
		}

		wwYaml, err := ww.ReadGlobalWorkWorkYaml()
		if err != nil {
			return err
		}

		log.Printf("projects: %+v", wwYaml.Projects)

		for _, project := range wwYaml.Projects {
			if project.Name == args.Project {
				cmd := exec.Command(args.Command)
				if err := cmd.Run(); err != nil {
					return fmt.Errorf("failed to execute command: %v", err)
				}
			}
		}

		return nil
	},
}
