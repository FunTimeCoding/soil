package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/filter_result"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/request"
)

func (s *Service) GetResource(
	x context.Context,
	clusterName string,
	q request.Get,
) (*filter_result.FilterResult, error) {
	return guarded(
		s,
		x,
		clusterName,
		func(c *cluster.Cluster) (*filter_result.FilterResult, error) {
			return getResourceOn(x, c, q)
		},
	)
}
