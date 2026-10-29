package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
)

func guarded[T any](
	s *Service,
	x context.Context,
	clusterName string,
	f func(*cluster.Cluster) (T, error),
) (T, error) {
	var result T
	c, e := s.ClusterByName(clusterName)

	if e != nil {
		return result, e
	}

	g := s.guards[clusterName].Attempt(
		x,
		func(context.Context) error {
			var h error
			result, h = f(c)

			return h
		},
	)

	return result, g
}
