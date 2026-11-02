package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/face"
)

func guarded[T any](
	s *Service,
	x context.Context,
	instance string,
	f func(face.Client) (T, error),
) (T, error) {
	var result T
	p, e := s.pool.Client(instance)

	if e != nil {
		return result, e
	}

	g := p.AlertmanagerGuard().Attempt(
		x,
		func(context.Context) error {
			var h error
			result, h = f(p.Alertmanager())

			return h
		},
	)

	return result, g
}
