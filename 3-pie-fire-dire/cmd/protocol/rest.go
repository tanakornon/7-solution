package protocol

import (
	"context"
	"net/http"
	_ "pie-fire-dire/docs" // Swagger docs
	"time"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title       Pie Fire Dire
// @description This is the API for the Pie Fire Dire application.
// @version     1.0

// @BasePath /
type RESTServer struct {
	app  *fiber.App
	port string
}

func NewRESTServer() *RESTServer {
	var (
		config  = app.config
		handler = app.handlers.rest
	)

	app := fiber.New()

	app.Get("/docs/*", fiberSwagger.WrapHandler)
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON("ok")
	})

	api := app.Group("/api")
	v1Group := api.Group("/v1")
	{
		public := v1Group
		{
			beef := public.Group("/beef")
			{
				h := handler.Beef
				beef.Get("/summary", h.GetSummary)
				beef.Get("/", h.Get)
			}
		}
	}

	return &RESTServer{
		app:  app,
		port: config.Server.RESTPort,
	}
}

func (s *RESTServer) Start() error {
	return s.app.Listen(":" + s.port)
}

func (s *RESTServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.app.ShutdownWithContext(ctx)
}
