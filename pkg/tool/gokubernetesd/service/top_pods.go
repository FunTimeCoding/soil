package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/request"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/response"
)

func (s *Service) TopPods(
	x context.Context,
	clusterName string,
	q request.Top,
) ([]response.PodMetrics, error) {
	return guarded(
		s,
		x,
		clusterName,
		func(c *cluster.Cluster) ([]response.PodMetrics, error) {
			return topPodsOn(x, c, q)
		},
	)
}
