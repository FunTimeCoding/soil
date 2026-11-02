package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/argocd_detail"
)

func (s *Service) ArgocdApplication(
	x context.Context,
	clusterName string,
	name string,
	unfiltered bool,
) (*argocd_detail.Detail, error) {
	return guarded(
		s,
		x,
		clusterName,
		func(c *cluster.Cluster) (*argocd_detail.Detail, error) {
			return argocdApplicationOn(x, c, name, unfiltered)
		},
	)
}
