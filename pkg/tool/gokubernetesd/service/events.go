package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/request"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/response"
)

func (s *Service) Events(
	x context.Context,
	clusterName string,
	q request.Events,
) ([]response.EventEntry, error) {
	return guarded(
		s,
		x,
		clusterName,
		func(c *cluster.Cluster) ([]response.EventEntry, error) {
			return s.eventsOn(x, c, clusterName, q)
		},
	)
}
