package matchers

import (
	"fmt"
	"github.com/a3herrera/api-test/container/logger"
	"github.com/a3herrera/api-test/service/search"
)

type itunesMatcher struct {
	uri string
}

func init() {
	matcher := itunesMatcher{
		uri: "https://itunes.apple.com/search",
	}
	search.Register("itunes", matcher)
}

type itunesItem struct {
	WrapperType string `json:"wrapperType"`
	Kind        string `json:"kind"`
	Artist      string `json:"artisName"`
	Collection  string `json:"collectionName"`
}
type itunesResult struct {
	Count   int          `json:"resultCount"`
	Results []itunesItem `json:"results"`
}

func (m itunesMatcher) Search(searchValue string, searchTerm string) ([]*search.Result, error) {
	logger.Log.Info("Start itunes search")

	var result itunesResult
	//TODO: Pendiente de manejar los terminos de busqueda
	fullUri := fmt.Sprintf("%s?term=jim+jones&limit25", m.uri)
	getRequest(fullUri, &result)
	// TODO: Ya tengo respuesta de itunes. Pendiente de ver como chutas manejar los resultados
	return nil, nil
}
