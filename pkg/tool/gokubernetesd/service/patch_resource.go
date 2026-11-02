package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/request"
)

func (s *Service) PatchResource(
	x context.Context,
	clusterName string,
	q request.Patch,
) error {
	return guardedError(
		s,
		x,
		clusterName,
		func(c *cluster.Cluster) error { return patchResourceOn(x, c, q) },
	)
}
