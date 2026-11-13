package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

func guarded[T any](
	x context.Context,
	s *Service,
	instance string,
	f func(context.Context, face.GitLabClient) (T, error),
) (T, error) {
	var result T
	p, e := s.pool.Client(instance)

	if e != nil {
		return result, e
	}

	return peer.Guarded(
		x,
		p.Guard(),
		func(y context.Context) (T, error) { return f(y, p.Client()) },
	)
}
