package protocol

import (
	"net/http"
	"pie-fire-dire/config"
	"pie-fire-dire/internal/adapter"
	"pie-fire-dire/internal/core/service"
	"pie-fire-dire/internal/handler/grpc"
	"pie-fire-dire/internal/handler/rest"
)

var app *application

type application struct {
	config   *config.AppConfig
	adapters *adapter.Adapters
	services *service.Services
	handlers *handlers
}

type handlers struct {
	grpc *grpc.Handlers
	rest *rest.Handlers
}

func init() {
	var (
		config     = config.Get()
		httpClient = http.DefaultClient

		adapters = adapter.New(adapter.Dependencies{
			Config:     config,
			HttpClient: httpClient,
		})

		services = service.New(service.Dependencies{
			Adapters: adapters,
		})

		handlers = &handlers{
			grpc: grpc.NewHandlers(grpc.Dependencies{
				Services: services,
			}),
			rest: rest.NewHandlers(rest.Dependencies{
				Services: services,
			}),
		}
	)

	app = &application{
		config:   config,
		adapters: adapters,
		services: services,
		handlers: handlers,
	}
}
