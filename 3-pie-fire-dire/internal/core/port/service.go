package port

import (
	"context"
	"pie-fire-dire/internal/core/domain"
)

type BeefService interface {
	Summary(ctx context.Context) (domain.BeefSummary, error)
}
