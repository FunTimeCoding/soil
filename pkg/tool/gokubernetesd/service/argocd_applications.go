package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/response"
)

func (s *Service) ArgocdApplications(
	x context.Context,
	clusterName string,
) ([]response.ArgocdSummary, error) {
	return guarded(
		s,
		x,
		clusterName,
		func(c *cluster.Cluster) ([]response.ArgocdSummary, error) {
			return argocdApplicationsOn(x, c)
		},
	)
}
