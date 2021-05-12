package matchers

import (
	"github.com/a3herrera/api-test/container/logger"
	"github.com/a3herrera/api-test/service/search"
)

type (
	itunesMatcher struct {
		uri  string
		path string
	}
	itunesItem struct {
		WrapperType string `json:"wrapperType"`
		Kind        string `json:"kind"`
		Artist      string `json:"artistName"`
		Collection  string `json:"collectionName"`
	}
	itunesResult struct {
		Count   int          `json:"resultCount"`
		Results []itunesItem `json:"results"`
	}
)

func init() {
	matcher := itunesMatcher{
		uri:  "https://itunes.apple.com",
		path: "/search",
	}
	search.Register("itunes", matcher)
}

func (m itunesMatcher) Search(searchValue string) (*search.Result, error) {
	logger.Log.Info("Start itunes search")

	var result itunesResult
	queryParams := map[string]string{
		"term": searchValue,
	}
	URI, err := GET(m.uri, m.path, queryParams, &result)
	if err != nil {
		return nil, err
	}
	searchResult := search.Result{
		URI:    URI,
		Exists: false,
	}

	if len(result.Results) > 0 {
		searchResult.Results = result.Results
		searchResult.Exists = true
	}

	return &searchResult, nil
}
