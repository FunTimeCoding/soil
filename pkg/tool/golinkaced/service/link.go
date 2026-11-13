package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/linkace/link"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) link(identifier int) (*link.Link, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*link.Link, error) {
			return s.client.Link(identifier)
		},
	)
}
