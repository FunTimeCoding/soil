package service

import "github.com/funtimecoding/soil/pkg/netbox/cluster"

func (s *Service) Clusters() ([]*cluster.Cluster, error) {
	return s.client.Clusters()
}
