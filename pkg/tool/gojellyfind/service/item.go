package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
)

func (s *Service) Item(identifier string) (*item.Item, error) {
	if identifier == "" {
		return nil, validation.New("id is required")
	}

	return s.source.Item(identifier)
}
