package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
)

func (s *Service) Episodes(seriesIdentifier string) ([]*item.Item, error) {
	if seriesIdentifier == "" {
		return nil, validation.New("series_id is required")
	}

	return s.source.Episodes(seriesIdentifier)
}
