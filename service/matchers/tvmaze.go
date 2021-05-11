package matchers

import (
	"github.com/a3herrera/api-test/container/logger"
	"github.com/a3herrera/api-test/service/search"
)

type tvMazeMatcher struct {
}

func init() {
	var matcher tvMazeMatcher
	search.Register("tv-maze", matcher)
}

func (m tvMazeMatcher) Search(searchValue string, searchTerm string) ([]*search.Result, error) {
	logger.Log.Info("Start tvMaze search")
	return nil, nil
}
