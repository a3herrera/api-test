package matchers

import (
	"github.com/a3herrera/api-test/container/logger"
	"github.com/a3herrera/api-test/service/search"
)

type crcindMatcher struct {
}

func init() {
	var matcher crcindMatcher
	search.Register("crcind", matcher)
}

func (m crcindMatcher) Search(searchValue string, searchTerm string) ([]*search.Result, error) {
	logger.Log.Info("Start crcind search")
	return nil, nil
}
