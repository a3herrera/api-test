package search

import (
	"github.com/a3herrera/api-test/container/logger"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchValue string, searchTerm string) []*Result {
	results := make(chan *Result)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(matchers))

	for _, matcher := range matchers {
		go func(matcher Matcher) {
			Match(matcher, searchValue, searchTerm, results)
			waitGroup.Done()
		}(matcher)
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()

	return retrieveResults(results)
}

func retrieveResults(results chan *Result) []*Result {
	resultsCompiler := make([]*Result, 0)
	for result := range results {
		resultsCompiler = append(resultsCompiler, result)
	}
	return resultsCompiler
}

func Register(searcherType string, matcher Matcher) {
	if _, exists := matchers[searcherType]; exists {
		logger.Log.Errorf("the matcher %s, is already registered", searcherType)
		return
	}
	logger.Log.Info("Registered ", searcherType, "matcher")
	matchers[searcherType] = matcher
}
