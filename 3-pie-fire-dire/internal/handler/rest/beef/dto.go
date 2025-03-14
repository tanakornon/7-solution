package beef

import (
	"pie-fire-dire/internal/core/domain"
)

type dto struct{}

func (t *dto) toSummaryResponse(in domain.BeefSummary) summary {
	return summary{
		Beef: in,
	}
}
