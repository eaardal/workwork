package commands

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"workwork/src/gui"
	"workwork/src/ww"
)

var LSCommand = &cli.Command{
	Name:      "ls",
	Usage:     "List all the registered URLs",
	UsageText: "Run like this: `ww ls`",
	Action: func(c *cli.Context) error {
		wwFile, err := ww.ReadWorkWorkFile()
		if err != nil {
			return err
		}

		hiWhite := color.New(color.FgHiWhite).SprintFunc()
		ui := gui.NewUserInterface()

		for key, value := range wwFile.Urls {
			ui.Write("%s\t%s\t", hiWhite(key), value)
		}

		for _, env := range wwFile.Environments {
			ui.Write("\n%s", gui.FgHiGreen(env.Name))

			for key, value := range env.Urls {
				ui.Write("%s\t%s\t", hiWhite(key), value)
			}
		}

		ui.MustFlush()
		return nil
	},
}

func debugStr(str string) {
	fmt.Printf("plain string: ")
	fmt.Printf("%s", str)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", str)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
	fmt.Printf("\n")
}
