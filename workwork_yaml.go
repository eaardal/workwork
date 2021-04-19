package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

const workWorkFileName = ".workwork.yaml"

type WorkWorkFile struct {
	Urls map[string]string `yaml:"urls"`
}

func absoluteWorkWorkFilePath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	filepath := path.Join(wd, workWorkFileName)
	return filepath, nil
}

func readWorkWorkFile() (*WorkWorkFile, error) {
	filepath, err := absoluteWorkWorkFilePath()
	if err != nil {
		return nil, err
	}

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var ww WorkWorkFile
	if err := yaml.Unmarshal(file, &ww); err != nil {
		return nil, err
	}

	return &ww, nil
}

func writeWorkWorkFile(ww *WorkWorkFile) error {
	filepath, err := absoluteWorkWorkFilePath()
	if err != nil {
		return err
	}

	wwYaml, err := yaml.Marshal(ww)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath, wwYaml, 0644); err != nil {
		return err
	}

	return nil
}
