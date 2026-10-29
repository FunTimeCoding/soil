package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/service/outcome"
)

func (s *Service) Collect(x context.Context) ([]*outcome.Outcome, error) {
	g, e := s.loadGazetteer(x)

	if e != nil {
		return nil, e
	}

	var result []*outcome.Outcome

	for _, c := range s.collectors {
		result = append(result, s.placeSource(x, c, g))
	}

	if s.lease != nil {
		result = append(result, s.sightLeases(x, g))
	}

	return result, nil
}
