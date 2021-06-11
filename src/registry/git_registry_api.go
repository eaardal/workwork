package registry

import "github.com/eaardal/workwork/src/ww"

type gitRegistryApi struct {
}

func NewGitRegistryApi() Api {
	return &gitRegistryApi{}
}

func (a gitRegistryApi) Search(query string) ([]SearchResult, error) {
	// - Git pull registry repo
	// - Look for directory where name contains query
	// - List all partially matching results
	panic("implement me")
}

func (a gitRegistryApi) Get(appName string) (*ww.WorkWorkYaml, error) {
	// - Git pull registry repo
	// - Look for directory matching appName
	// - Copy .workwork.yaml from registry repo to app repo
	panic("implement me")
}

func (a gitRegistryApi) Publish(wwFile ww.WorkWorkYaml) error {
	// - Git pull registry repo
	// - Copy .workwork.yaml to git repository on disk
	// - Git commit and tag registry repo
	// - Git push to remote
	return nil
}
