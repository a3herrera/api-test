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

func (hh SearchHandler) search(w http.ResponseWriter, _ *http.Request) error {
	_ = json.NewEncoder(w).Encode("pong")
	return nil
}
