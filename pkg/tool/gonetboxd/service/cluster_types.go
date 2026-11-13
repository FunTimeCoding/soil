package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_type"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) ClusterTypes() ([]*cluster_type.Type, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*cluster_type.Type, error) {
			return s.client.ClusterTypes()
		},
	)
}
