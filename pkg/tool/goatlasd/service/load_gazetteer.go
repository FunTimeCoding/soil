package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
)

func (s *Service) loadGazetteer(x context.Context) (*gazetteer.Gazetteer, error) {
	var result *gazetteer.Gazetteer
	e := s.netboxGuard.Attempt(
		x,
		func(y context.Context) error {
			var f error
			result, f = gazetteer.Load(y, s.netbox)

			return f
		},
	)

	if e != nil {
		return nil, e
	}

	return result, nil
}
