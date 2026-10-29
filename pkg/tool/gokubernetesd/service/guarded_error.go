package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
)

func guardedError(
	s *Service,
	x context.Context,
	clusterName string,
	f func(*cluster.Cluster) error,
) error {
	_, e := guarded(
		s,
		x,
		clusterName,
		func(c *cluster.Cluster) (struct{}, error) {
			return struct{}{}, f(c)
		},
	)

	return e
}
