package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/describe_result"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
)

func (s *Service) DescribeResource(
	x context.Context,
	clusterName string,
	q request.Describe,
) (*describe_result.Result, error) {
	return guarded(
		s,
		x,
		clusterName,
		func(c *cluster.Cluster) (*describe_result.Result, error) {
			return describeResourceOn(x, c, q)
		},
	)
}
