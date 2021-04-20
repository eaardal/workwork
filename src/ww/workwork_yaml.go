package ww

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

const workWorkFileName = ".workwork.yaml"
const Global = "global"

type Environment struct {
	Name            string `yaml:"name"`
	EnvironmentUrls Urls   `yaml:"urls"`
}

func NewEnvironment(name string, urls map[string]string) Environment {
	return Environment{
		Name:            name,
		EnvironmentUrls: urls,
	}
}

type WorkWorkFile struct {
	GlobalUrls   Urls          `yaml:"urls"`
	Environments []Environment `yaml:"environments"`
}

func (f WorkWorkFile) GetEnvironment(env string) (*Environment, error) {
	for _, environment := range f.Environments {
		if env == environment.Name {
			return &environment, nil
		}
	}
	return nil, fmt.Errorf("no environment named '%s'", env)
}

func (f WorkWorkFile) GetUrls(environmentName string) (map[string]string, error) {
	if environmentName == "" || environmentName == Global {
		return f.GlobalUrls, nil
	}

	env, err := f.GetEnvironment(environmentName)
	if err != nil {
		return nil, err
	}

	return env.EnvironmentUrls, nil
}

func ReadWorkWorkFile() (*WorkWorkFile, error) {
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

func WriteWorkWorkFile(ww *WorkWorkFile) error {
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

func absoluteWorkWorkFilePath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	filepath := path.Join(wd, workWorkFileName)
	return filepath, nil
}
