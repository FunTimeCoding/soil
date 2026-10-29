package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
)

func (s *Service) Tracks(albumIdentifier string) ([]*item.Item, error) {
	if albumIdentifier == "" {
		return nil, validation.New("album_id is required")
	}

	return s.source.Tracks(albumIdentifier)
}
