package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/band/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

func guarded[T any](
	s *Service,
	instance string,
	f func(face.BandClient) (T, error),
) (T, error) {
	var result T
	p, e := s.pool.Client(instance)

	if e != nil {
		return result, e
	}

	return peer.Guarded(
		context.Background(),
		p.Guard(),
		func(context.Context) (T, error) { return f(p.Client()) },
	)
}
