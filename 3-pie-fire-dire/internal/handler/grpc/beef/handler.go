package beef

import (
	"context"
	"pie-fire-dire/generated/grpc/beefpb"
	"pie-fire-dire/internal/core/port"
)

type Dependencies struct {
	Service port.BeefService
}

type Handler struct {
	beefpb.UnimplementedBeefServiceServer
	service port.BeefService
	dto     dto
}

func NewHandler(deps Dependencies) *Handler {
	return &Handler{
		service: deps.Service,
		dto:     dto{},
	}
}

func (s *Handler) Get(ctx context.Context, req *beefpb.GetRequest) (*beefpb.GetResponse, error) {
	return &beefpb.GetResponse{Message: "Not Implemented"}, nil
}

func (s *Handler) GetSummary(ctx context.Context, req *beefpb.GetSummaryRequest) (*beefpb.GetSummaryResponse, error) {
	summary, err := s.service.Summary(ctx)
	if err != nil {
		return nil, err
	}

	out := s.dto.toSummaryResponse(summary)
	return &out, nil
}
