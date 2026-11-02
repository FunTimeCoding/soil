package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/prometheus/rule"
)

func (s *Service) Rules(
	x context.Context,
	instance string,
) ([]*rule.Rule, error) {
	return guarded(
		s,
		x,
		instance,
		func(c face.Client) ([]*rule.Rule, error) {
			v, e := c.Rules()

			if e != nil {
				return nil, e
			}

			return v.Get(), nil
		},
	)
}
