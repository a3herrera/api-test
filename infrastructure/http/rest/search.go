package rest

import (
	"encoding/json"
	"github.com/a3herrera/api-test/service"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type SearchHandler struct {
	service service.SearchService
}

func NewSearchHandler(searchService service.SearchService) SearchHandler {
	return SearchHandler{
		service: searchService,
	}
}
func (hh SearchHandler) NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Method(http.MethodGet, "/", handler(hh.search))
	return r
}

func (hh SearchHandler) search(w http.ResponseWriter, r *http.Request) error {
	searchTerm := r.URL.Query().Get("value")
	if searchTerm == "" {
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"results": make([]interface{}, 0)})
		return nil
	}
	result := hh.service.Search(searchTerm)
	_ = json.NewEncoder(w).Encode(result)
	return nil
}
