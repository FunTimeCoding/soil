package service

import "github.com/funtimecoding/soil/pkg/netbox/cluster_type"

func (s *Service) CreateClusterType(name string) (*cluster_type.Type, error) {
	return s.client.CreateClusterType(name)
}
