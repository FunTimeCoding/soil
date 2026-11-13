package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/linkace/list"
	"github.com/funtimecoding/soil/pkg/linkace/page"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) ListsPage(p int) (*page.Page[*list.List], error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*page.Page[*list.List], error) {
			return s.client.ListsPage(p)
		},
	)
}
