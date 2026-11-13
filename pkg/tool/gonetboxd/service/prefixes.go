package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Prefixes() ([]*prefix.Prefix, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*prefix.Prefix, error) {
			return s.client.Prefixes()
		},
	)
}
