package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os/exec"
	"runtime"
	"workwork/src/gui"
	"workwork/src/ww"
)

var GoToCommand = &cli.Command{
	Name:      "goto",
	Usage:     "Open the URL for the given key using the default browser",
	UsageText: "Any URLs listed by `ww ls` can be opened with this command: `ww goto {key}`. If you're on Mac, the `open` executable will be used. On Linux, it'll check for [open, xdg-open] in that order (more advanced Linux support probably needed). On Windows, the `start` executable will be used.",
	Action: func(c *cli.Context) error {
		ui := gui.NewUserInterface()

		args, err := ParseAndValidateGoToCommandArgs(c)
		if err != nil {
			return err
		}

		wwFile, err := ww.ReadWorkWorkYaml()
		if err != nil {
			return err
		}

		urls, err := wwFile.GetUrls(args.Environment)
		if err != nil {
			return err
		}

		foundMatch := false

		for urlKey, url := range urls {
			if args.UrlKey == urlKey {
				var err error

				switch runtime.GOOS {
				case "windows":
					err = openOnWindows(url)
					break
				case "darwin":
					err = openOnMac(url)
					break
				case "linux":
					err = openOnLinux(url)
					break
				default:
					return fmt.Errorf("unsupported GOOS '%s'", runtime.GOOS)
				}

				if err != nil {
					return err
				}

				foundMatch = true
				break
			}
		}

		if !foundMatch {
			return fmt.Errorf("found no url with key '%s'", args.UrlKey)
		}

		ui.MustFlush()
		return nil
	},
}

func openOnMac(url string) error {
	return exec.Command("open", url).Run()
}

func openOnWindows(url string) error {
	windowTitle := ""
	return exec.Command("start", windowTitle, url).Run()
}

func openOnLinux(url string) error {
	path, err := exec.LookPath("open")
	if err == nil && path != "" {
		return exec.Command(path, url).Run()
	}

	path, err = exec.LookPath("xdg-open")
	if err == nil && path != "" {
		return exec.Command(path, url).Run()
	}

	return fmt.Errorf("unable to open '%s'. Looked for executables ['open', 'xdg-open'] but found none", url)
}
