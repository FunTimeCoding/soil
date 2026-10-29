package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/response"
)

func (s *Service) ListResources(
	x context.Context,
	clusterName string,
	q request.List,
) ([]response.ResourceSummary, error) {
	return guarded(
		s,
		x,
		clusterName,
		func(c *cluster.Cluster) ([]response.ResourceSummary, error) {
			return listResourcesOn(x, c, q)
		},
	)
}
