package matchers

import (
	"github.com/a3herrera/api-test/container/logger"
	"github.com/a3herrera/api-test/service/search"
)

type itunesMatcher struct {
}

func init() {
	var matcher itunesMatcher
	search.Register("itunes", matcher)
}

func (m itunesMatcher) Search(searchValue string, searchTerm string) ([]*search.Result, error) {
	logger.Log.Info("Start itunes search")

	return nil, nil
}
