package ww

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

type Repository struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

type Project struct {
	Name  string       `yaml:"name"`
	Repos []Repository `yaml:"repositories"`
}

type GlobalWorkWorkYaml struct {
	Projects []Project `yaml:"projects"`
}

func ReadGlobalWorkWorkYaml() (*GlobalWorkWorkYaml, error) {
	filepath, err := AbsoluteGlobalWorkWorkYamlFilePath()
	if err != nil {
		return nil, err
	}

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var ww GlobalWorkWorkYaml
	if err := yaml.Unmarshal(file, &ww); err != nil {
		return nil, err
	}

	return &ww, nil
}

func WriteGlobalWorkWorkYaml(ww *GlobalWorkWorkYaml) error {
	filepath, err := AbsoluteGlobalWorkWorkYamlFilePath()
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

func AbsoluteGlobalWorkWorkYamlFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filepath := path.Join(home, workWorkFileName)
	return filepath, nil
}
