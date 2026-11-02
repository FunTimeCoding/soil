package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/types/response"
)

func (s *Service) Certificates(
	x context.Context,
	clusterName string,
) ([]response.CertificateSummary, error) {
	return guarded(
		s,
		x,
		clusterName,
		func(c *cluster.Cluster) ([]response.CertificateSummary, error) {
			return certificatesOn(x, c)
		},
	)
}
