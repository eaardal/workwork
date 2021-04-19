package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os/exec"
	"runtime"
)

var gotoCmd = &cli.Command{
	Name:  "goto",
	Usage: "Go to the URL for the given key, using the OS default browser",
	Action: func(c *cli.Context) error {
		ui := newUserInterface()

		key := c.Args().Get(0)

		if !isValidKey(key) {
			return fmt.Errorf("invalid key '%s'", key)
		}

		ww, err := readWorkWorkFile()
		if err != nil {
			return err
		}

		foundMatch := false
		for wwkey, url := range ww.Urls {
			if key == wwkey {
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
			return fmt.Errorf("found no url with key '%s'", key)
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
