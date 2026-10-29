package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/response"
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
