package rest

import (
	"pie-fire-dire/internal/core/service"
	"pie-fire-dire/internal/handler/rest/beef"
)

type Dependencies struct {
	Services *service.Services
}

type Handlers struct {
	Beef *beef.Handler
}

func NewHandlers(deps Dependencies) *Handlers {
	return &Handlers{
		Beef: beef.NewHandler(beef.Dependencies{
			Service: deps.Services.BaconService,
		}),
	}
}
