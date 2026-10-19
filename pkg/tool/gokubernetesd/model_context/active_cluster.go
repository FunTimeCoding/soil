package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
)

func (s *Server) activeCluster(x context.Context) (*cluster.Cluster, error) {
	name, e := s.activeClusterName(x)

	if e != nil {
		return nil, e
	}

	return s.service.ClusterByName(name)
}
