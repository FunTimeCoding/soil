package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Item(identifier string) (*item.Item, error) {
	if identifier == "" {
		return nil, validation.New("id is required")
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*item.Item, error) {
			return s.client.Item(identifier)
		},
	)
}
