package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/response"
)

func (s *Service) TopNodes(
	x context.Context,
	clusterName string,
) ([]response.NodeMetrics, error) {
	return guarded(
		s,
		x,
		clusterName,
		func(c *cluster.Cluster) ([]response.NodeMetrics, error) {
			return topNodesOn(x, c)
		},
	)
}
