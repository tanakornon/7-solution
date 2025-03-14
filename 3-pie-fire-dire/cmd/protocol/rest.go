package protocol

import (
	"net/http"
	_ "pie-fire-dire/docs" // Swagger docs

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title       Pie Fire Dire
// @description This is the API for the Pie Fire Dire application.
// @version     1.0

// @BasePath /

func ServeREST() error {
	var (
		config   = app.config
		handlers = app.handlers.rest
	)

	router := fiber.New()

	router.Get("/docs/*", fiberSwagger.WrapHandler)
	router.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON("ok")
	})

	api := router.Group("/api")
	v1Group := api.Group("/v1")
	{
		public := v1Group
		{
			beef := public.Group("/beef")
			{
				h := handlers.Beef
				beef.Get("/summary", h.GetSummary)
				beef.Get("/", h.Get)
			}
		}
	}

	return router.Listen(":" + config.Server.RESTPort)
}
