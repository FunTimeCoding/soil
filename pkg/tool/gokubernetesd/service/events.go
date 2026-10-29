package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/response"
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
