package search

import "github.com/a3herrera/api-test/container/logger"

type Result struct {
	URI     string      `json:"uri"`
	Exists  bool        `json:"-"`
	Results interface{} `json:"results"`
}

type Matcher interface {
	Search(searchValue string) (*Result, error)
}

func Match(matcher Matcher, searchValue string, results chan<- *Result) {
	searchResults, err := matcher.Search(searchValue)
	if err != nil {
		logger.Log.Errorf("Fail in the execution in the matcher, error: %v", err)
		return
	}
	results <- searchResults
}
