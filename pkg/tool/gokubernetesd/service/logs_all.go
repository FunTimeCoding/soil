package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/request"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/response"
)

func (s *Service) LogsAll(
	x context.Context,
	clusterName string,
	q request.Logs,
) ([]response.PodLog, error) {
	return guarded(
		s,
		x,
		clusterName,
		func(c *cluster.Cluster) ([]response.PodLog, error) {
			return logsAllOn(x, c, q)
		},
	)
}
