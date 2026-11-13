package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/opnsense/lease"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Leases(phrase string) ([]*lease.Lease, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*lease.Lease, error) {
			return s.client.Leases(phrase)
		},
	)
}
