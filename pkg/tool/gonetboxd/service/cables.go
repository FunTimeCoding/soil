package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/cable"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Cables() ([]*cable.Cable, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*cable.Cable, error) {
			return s.client.Cables()
		},
	)
}
