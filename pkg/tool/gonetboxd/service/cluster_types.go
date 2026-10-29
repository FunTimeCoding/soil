package service

import "github.com/funtimecoding/soil/pkg/netbox/cluster_type"

func (s *Service) ClusterTypes() ([]*cluster_type.Type, error) {
	return s.client.ClusterTypes()
}
