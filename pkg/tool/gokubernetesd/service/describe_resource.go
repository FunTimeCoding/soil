package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/describe_result"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/request"
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
