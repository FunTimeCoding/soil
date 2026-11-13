package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/habitica/gear"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) UserGear() (*gear.Gear, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*gear.Gear, error) {
			return s.client.UserGear()
		},
	)
}
