package matchers

import (
	"github.com/a3herrera/api-test/container/logger"
	"github.com/a3herrera/api-test/service/search"
)

type (
	tvMazeMatcher struct {
		uri  string
		path string
	}

	tvMazeShow struct {
		ID       int      `json:"id,omitempty"`
		URL      string   `json:"url,omitempty"`
		Name     string   `json:"name,omitempty"`
		Type     string   `json:"type,omitempty"`
		Language string   `json:"language,omitempty"`
		Genres   []string `json:"genres,omitempty"`
		Status   string   `json:"status,omitempty"`
	}

	tvMazeItem struct {
		Score    float64     `json:"score,omitempty"`
		Show     *tvMazeShow `json:"show,omitempty"`
		ID       int         `json:"id,omitempty"`
		URL      string      `json:"url,omitempty"`
		Name     string      `json:"name,omitempty"`
		Type     string      `json:"type,omitempty"`
		Language string      `json:"language,omitempty"`
		Genres   []string    `json:"genres,omitempty"`
		Status   string      `json:"status,omitempty"`
	}
)

func init() {
	matcher := tvMazeMatcher{
		uri:  "http://api.tvmaze.com/search?q=",
		path: "/shows",
	}
	search.Register("tv-maze", matcher)
}

func (m tvMazeMatcher) Search(searchValue string) (*search.Result, error) {
	logger.Log.Info("Start tvMaze search")
	results := make([]tvMazeItem, 0)

	queryParams := map[string]string{
		"q": searchValue,
	}
	URI, err := GET(m.uri, m.path, queryParams, &results)
	if err != nil {
		return nil, err
	}

	searchResult := search.Result{
		URI:    URI,
		Exists: false,
	}

	if len(results) > 0 {
		searchResult.Results = results
		searchResult.Exists = true
	}
	return &searchResult, nil
}
