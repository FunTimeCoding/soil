package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/request"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/response"
)

func (s *Service) TopPodContainers(
	x context.Context,
	clusterName string,
	q request.Top,
) ([]response.ContainerMetrics, error) {
	return guarded(
		s,
		x,
		clusterName,
		func(c *cluster.Cluster) ([]response.ContainerMetrics, error) {
			return topPodContainersOn(x, c, q)
		},
	)
}
