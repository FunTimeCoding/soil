package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
)

func (s *Service) DeleteResource(
	x context.Context,
	clusterName string,
	resourceType string,
	name string,
	namespace string,
) error {
	return guardedError(
		s,
		x,
		clusterName,
		func(c *cluster.Cluster) error {
			return deleteResourceOn(x, c, resourceType, name, namespace)
		},
	)
}
