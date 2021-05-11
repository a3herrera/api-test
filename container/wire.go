// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package container

import (
	"github.com/a3herrera/api-test/infrastructure/http"
	"github.com/a3herrera/api-test/infrastructure/http/rest"
	"github.com/a3herrera/api-test/service"
	"github.com/google/wire"
)

var servicesSet = wire.NewSet(
	service.NewApisService,
)

var handlerSet = wire.NewSet(
	rest.NewHealthHandler,
	http.NewServerHandlers,
)

func NewServer() (http.Server, error) {
	wire.Build(handlerSet, http.New)
	return http.Server{}, nil
}
