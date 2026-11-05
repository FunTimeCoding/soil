package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
)

func (s *Service) Item(identifier string) (*item.Item, error) {
	if identifier == "" {
		return nil, validation.New("id is required")
	}

	return guarded(
		s.guard,
		func() (*item.Item, error) {
			return s.client.Item(identifier)
		},
	)
}
