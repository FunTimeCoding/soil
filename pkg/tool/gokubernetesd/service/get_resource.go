package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/resource/filter_result"
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
