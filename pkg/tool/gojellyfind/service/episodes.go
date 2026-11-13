package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Episodes(seriesIdentifier string) ([]*item.Item, error) {
	if seriesIdentifier == "" {
		return nil, validation.New("series_id is required")
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*item.Item, error) {
			return s.client.Episodes(seriesIdentifier)
		},
	)
}
