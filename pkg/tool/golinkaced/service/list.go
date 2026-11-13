package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/linkace/list"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) List(name string) (*list.List, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*list.List, error) { return s.client.List(name) },
	)
}
