package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/opnsense/pool"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Pools(phrase string) ([]*pool.Pool, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*pool.Pool, error) {
			return s.client.Pools(phrase)
		},
	)
}
