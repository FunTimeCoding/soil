package service

import "github.com/funtimecoding/soil/pkg/netbox/cluster"

func (s *Service) CreateCluster(
	name string,
	clusterType string,
	site string,
) (*cluster.Cluster, error) {
	t, e := s.client.ClusterTypeByName(clusterType)

	if e != nil {
		return nil, e
	}

	i, e := s.client.SiteByName(site)

	if e != nil {
		return nil, e
	}

	return s.client.CreateCluster(name, t, i)
}
