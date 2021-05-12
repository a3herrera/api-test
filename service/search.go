package service

import (
	"github.com/a3herrera/api-test/service/search"
)

type SearchService struct {
}

func NewSearchService() SearchService {
	return SearchService{}
}

func (as SearchService) Search(searchValue string) interface{} {
	searchResult := search.Run(searchValue)
	results := make([]*search.Result, 0)
	for _, item := range searchResult {
		if item.Exists {
			results = append(results, item)
		}
	}

	return map[string][]*search.Result{"results": results}
}
