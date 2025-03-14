package service

import (
	"pie-fire-dire/internal/adapter"
	"pie-fire-dire/internal/core/port"
	"pie-fire-dire/internal/core/service/beef"
)

type Dependencies struct {
	Adapters *adapter.Adapters
}

type Services struct {
	BaconService port.BeefService
}

func New(deps Dependencies) *Services {
	return &Services{
		BaconService: beef.NewService(beef.Dependencies{
			BaconIpsumAPI: deps.Adapters.BaconIpsumAPI,
		}),
	}
}
