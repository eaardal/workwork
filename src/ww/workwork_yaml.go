package ww

import (
	"bufio"
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

type WorkWorkYaml struct {
	GlobalUrls   Urls          `yaml:"urls"`
	Environments []Environment `yaml:"environments"`
}

func (f WorkWorkYaml) GetEnvironment(env string) (*Environment, error) {
	for _, environment := range f.Environments {
		if env == environment.Name {
			return &environment, nil
		}
	}
	return nil, fmt.Errorf("no environment named '%s'", env)
}

func (f WorkWorkYaml) GetUrls(environmentName string) (map[string]string, error) {
	if environmentName == "" || environmentName == Global {
		return f.GlobalUrls, nil
	}

	env, err := f.GetEnvironment(environmentName)
	if err != nil {
		return nil, err
	}

	return env.EnvironmentUrls, nil
}

func ReadWorkWorkYaml(rootDir string) (*WorkWorkYaml, error) {
	filepath, err := AbsoluteWorkWorkYamlFilePath(rootDir)
	if err != nil {
		return nil, err
	}

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var ww WorkWorkYaml
	if err := yaml.Unmarshal(file, &ww); err != nil {
		return nil, err
	}

	return &ww, nil
}

func WriteWorkWorkYaml(rootDir string, ww *WorkWorkYaml) error {
	filepath, err := AbsoluteWorkWorkYamlFilePath(rootDir)
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

	//if err := prependHeaderToYamlFile(filepath); err != nil {
	//	return err
	//}

	return nil
}

func AbsoluteWorkWorkYamlFilePath(rootDir string) (string, error) {
	filepath := path.Join(rootDir, workWorkFileName)
	return filepath, nil
}

func prependHeaderToYamlFile(filepath string) (e error) {
	headerLines := `
# This file was made by Workwork, a CLI tool for listing and opening URLs for common software development concerns.
# Read more at: https://github.com/eaardal/workwork
`

	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer func() {
		if err := f.Close(); err != nil {
			e = err
		}
	}()

	writer := bufio.NewWriter(f)
	if _, err := writer.WriteString(headerLines); err != nil {
		return err
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	return nil
}
