package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/linkace/page"
	"github.com/funtimecoding/soil/pkg/linkace/tag"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) TagsPage(p int) (*page.Page[*tag.Tag], error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*page.Page[*tag.Tag], error) {
			return s.client.TagsPage(p)
		},
	)
}
