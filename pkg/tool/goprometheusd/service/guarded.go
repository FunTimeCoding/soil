package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
)

func guarded[T any](
	s *Service,
	x context.Context,
	instance string,
	f func(face.PrometheusClient) (T, error),
) (T, error) {
	var result T
	p, e := s.pool.Client(instance)

	if e != nil {
		return result, e
	}

	return peer.Guarded(
		x,
		p.Guard(),
		func(context.Context) (T, error) { return f(p.Client()) },
	)
}
