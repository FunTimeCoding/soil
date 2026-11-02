package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/request"
)

func (s *Service) Logs(
	x context.Context,
	clusterName string,
	q request.Logs,
) (string, error) {
	return guarded(
		s,
		x,
		clusterName,
		func(c *cluster.Cluster) (string, error) { return logsOn(x, c, q) },
	)
}
