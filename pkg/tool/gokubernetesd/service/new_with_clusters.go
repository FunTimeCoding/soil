package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store"
)

func NewWithClusters(
	s *store.Store,
	clusters map[string]*cluster.Cluster,
) *Service {
	return &Service{clusters: clusters, store: s}
}
