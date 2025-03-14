package beef

import (
	"pie-fire-dire/generated/grpc/beefpb"
	"pie-fire-dire/internal/core/domain"
)

type dto struct{}

func (t *dto) toSummaryResponse(in domain.BeefSummary) beefpb.GetSummaryResponse {
	out := make(map[string]int64)
	for k, v := range in {
		out[k] = int64(v)
	}

	return beefpb.GetSummaryResponse{
		Beef: out,
	}
}
