package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
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

	return peer.Guarded(
		x,
		s.guards[clusterName],
		func(context.Context) (T, error) { return f(c) },
	)
}
