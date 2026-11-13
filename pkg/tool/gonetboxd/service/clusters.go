package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Clusters() ([]*cluster.Cluster, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*cluster.Cluster, error) {
			return s.client.Clusters()
		},
	)
}
