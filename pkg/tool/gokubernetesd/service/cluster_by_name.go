package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
)

func (s *Service) ClusterByName(name string) (*cluster.Cluster, error) {
	c, okay := s.clusters[name]

	if !okay {
		return nil, not_found.New("cluster", name)
	}

	return c, nil
}
