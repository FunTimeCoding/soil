package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
)

func guarded[T any](
	s *Service,
	call func() (T, error),
) (T, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (T, error) {
			return call()
		},
	)
}
