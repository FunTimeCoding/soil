package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/bookmark"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Bookmarks() ([]*bookmark.Bookmark, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*bookmark.Bookmark, error) {
			return s.client.Bookmarks()
		},
	)
}
