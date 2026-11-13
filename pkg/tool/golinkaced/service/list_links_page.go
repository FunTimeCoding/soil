package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/linkace/link"
	"github.com/funtimecoding/soil/pkg/linkace/page"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) ListLinksPage(
	identifier int,
	p int,
) (*page.Page[*link.Link], error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*page.Page[*link.Link], error) {
			return s.client.LinksByListPage(identifier, p)
		},
	)
}
