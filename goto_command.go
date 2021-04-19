package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os/exec"
	"runtime"
)

var gotoCommand = &cli.Command{
	Name:      "goto",
	Usage:     "Open the URL for the given key using the default browser",
	UsageText: "Any URLs listed by `ww ls` can be opened with this command: `ww goto {key}`. If you're on Mac, the `open` executable will be used. On Linux, it'll check for [open, xdg-open] in that order (more advanced Linux support probably needed). On Windows, the `start` executable will be used.",
	Action: func(c *cli.Context) error {
		ui := newUserInterface()

		args := c.Args()

		var keyArg string
		var envArg string

		if args.Len() == 2 {
			envArg = args.Get(0)
			keyArg = args.Get(1)
		} else {
			keyArg = c.Args().Get(0)
			envArg = ""
		}

		if !isValidKey(keyArg) {
			return fmt.Errorf("invalid key '%s'", keyArg)
		}

		ww, err := readWorkWorkFile()
		if err != nil {
			return err
		}

		urls := ww.Urls
		if envArg != "" {
			env, err := ww.GetEnvironment(envArg)
			if err != nil {
				return err
			}
			urls = env.Urls
		}

		foundMatch := false
		for urlKey, url := range urls {
			if keyArg == urlKey {
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
			return fmt.Errorf("found no url with key '%s'", keyArg)
		}

		ui.mustFlush()
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
