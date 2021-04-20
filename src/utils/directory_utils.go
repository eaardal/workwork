package utils

import "os"

func CreateDirectoryIfNotExists(dir string) error {
	mode := os.ModeDir
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, mode); err != nil {
			return err
		}
	}
	return nil
}
