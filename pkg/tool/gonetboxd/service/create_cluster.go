package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreateCluster(
	name string,
	clusterType string,
	site string,
) (*cluster.Cluster, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*cluster.Cluster, error) {
			t, e := s.client.ClusterType(clusterType)

			if e != nil {
				return nil, e
			}

			i, e := s.client.Site(site)

			if e != nil {
				return nil, e
			}

			return s.client.CreateCluster(name, t, i)
		},
	)
}
