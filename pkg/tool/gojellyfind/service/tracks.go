package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Tracks(albumIdentifier string) ([]*item.Item, error) {
	if albumIdentifier == "" {
		return nil, validation.New("album_id is required")
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*item.Item, error) {
			return s.client.Tracks(albumIdentifier)
		},
	)
}
