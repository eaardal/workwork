package registry

import "github.com/eaardal/workwork/src/ww"

type SearchResult struct {
}

type Api interface {
	Search(query string) ([]SearchResult, error)
	Get(appName string) (*ww.WorkWorkYaml, error)
	Publish(wwFile ww.WorkWorkYaml) error
}
