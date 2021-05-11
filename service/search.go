package service

import "github.com/a3herrera/api-test/service/search"

type SearchService struct {
}

func NewSearchService() SearchService {
	return SearchService{}
}

func (as SearchService) Search(searchValue string, searchTerm string) interface{} {
	searchResult := search.Run(searchValue, searchTerm)
	results := make(map[string]interface{})
	results["result"] = searchResult
	return results
}
