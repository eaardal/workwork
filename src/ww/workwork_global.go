package ww

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"workwork/src/utils"
)

const globalWorkWorkDirName = ".workwork"

type WorkWorkGlobal struct {
	Links map[string]string `yaml:"links"`
}

func EnsureGlobalDirectoryExists() (string, error) {
	dir, err := GlobalDirectoryPath()
	if err != nil {
		return "", err
	}
	return dir, utils.CreateDirectoryIfNotExists(dir)
}

func GlobalDirectoryPath() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return globalDirectoryPathOnWindows()
	case "darwin":
		return globalDirectoryPathOnMac()
	case "linux":
		return globalDirectoryPathOnLinux()
	default:
		return "", fmt.Errorf("unsupported GOOS '%s'", runtime.GOOS)
	}
}

func globalDirectoryPathOnWindows() (string, error) {
	return "", nil // TODO
}

func globalDirectoryPathOnMac() (string, error) {
	home := os.Getenv("HOME")
	return path.Join(home, globalWorkWorkDirName), nil
}

func globalDirectoryPathOnLinux() (string, error) {
	home := os.Getenv("HOME")
	return path.Join(home, globalWorkWorkDirName), nil
}
