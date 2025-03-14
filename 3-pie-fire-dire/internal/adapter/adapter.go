package adapter

import (
	"net/http"
	"pie-fire-dire/config"
	"pie-fire-dire/internal/adapter/bacon"
	"pie-fire-dire/internal/core/port"
)

type Dependencies struct {
	Config     *config.AppConfig
	HttpClient *http.Client
}

type Adapters struct {
	BaconIpsumAPI port.BaconIpsumAPI
}

func New(deps Dependencies) *Adapters {
	return &Adapters{
		BaconIpsumAPI: bacon.New(bacon.Dependencies{
			Client: deps.HttpClient,
			URL:    deps.Config.BaconIpsum.URL,
		}),
	}
}
