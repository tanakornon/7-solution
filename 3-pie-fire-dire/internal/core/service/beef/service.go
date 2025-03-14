package beef

import (
	"context"
	"pie-fire-dire/internal/core/domain"
	"pie-fire-dire/internal/core/port"
	"regexp"
)

var _ port.BeefService = (*Service)(nil)

type Dependencies struct {
	BaconIpsumAPI port.BaconIpsumAPI
}

type Service struct {
	baconIpsumAPI port.BaconIpsumAPI
	regex         *regexp.Regexp
}

func NewService(deps Dependencies) *Service {
	return &Service{
		baconIpsumAPI: deps.BaconIpsumAPI,
		regex:         regexp.MustCompile(`\w+(?:-\w+)*`),
	}
}

func (s *Service) Summary(ctx context.Context) (domain.BeefSummary, error) {
	text, err := s.baconIpsumAPI.GetText(ctx)
	if err != nil {
		return nil, err
	}

	tokens := s.regex.FindAllString(text, -1)

	summary := make(map[string]int, len(tokens))
	for _, token := range tokens {
		summary[token]++
	}

	return summary, nil
}
