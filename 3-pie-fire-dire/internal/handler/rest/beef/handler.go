package beef

import (
	"net/http"
	"pie-fire-dire/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type Dependencies struct {
	Service port.BeefService
}

type Handler struct {
	service port.BeefService
	dto     dto
}

func NewHandler(deps Dependencies) *Handler {
	return &Handler{
		service: deps.Service,
		dto:     dto{},
	}
}

// Get fetches beef information
// @Summary Get beef details
// @Tags    Beef
// @Produce json
// @Param   version path     string true "App Version" Enums(v1)
// @Failure 501     {object} ErrorResponse
// @Router  /api/{version}/beef [GET]
func (h *Handler) Get(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusNotImplemented).JSON(ErrorResponse{"Not Implement"})
}

// GetSummary fetches the summary of beef
// @Summary Get beef summary
// @Tags    Beef
// @Produce json
// @Param   version path     string true "App Version" Enums(v1)
// @Success 200     {object} GetSummaryResponse
// @Failure 500     {object} ErrorResponse
// @Router  /api/{version}/beef/summary [GET]
func (h *Handler) GetSummary(ctx *fiber.Ctx) error {
	summary, err := h.service.Summary(ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(ErrorResponse{err.Error()})
	}

	out := h.dto.toSummaryResponse(summary)
	return ctx.Status(http.StatusOK).JSON(out)
}
