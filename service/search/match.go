package search

import "github.com/a3herrera/api-test/container/logger"

type Result struct {
	URI string `json:"uri"`
	//TODO: Pendiente de definir como manejar las respuestas de los diferentes sitios
}

type Matcher interface {
	Search(searchValue string, searchTerm string) ([]*Result, error)
}

func Match(matcher Matcher, searchValue string, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(searchValue, searchTerm)
	if err != nil {
		logger.Log.Errorf("Fail in the execution in the matcher, error: %v", err)
	}
	for _, result := range searchResults {
		results <- result
	}
}
