package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/certificate_detail"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
)

func (s *Service) Certificate(
	x context.Context,
	clusterName string,
	name string,
	namespace string,
	unfiltered bool,
) (*certificate_detail.Detail, error) {
	return guarded(
		s,
		x,
		clusterName,
		func(c *cluster.Cluster) (*certificate_detail.Detail, error) {
			return certificateOn(x, c, name, namespace, unfiltered)
		},
	)
}
