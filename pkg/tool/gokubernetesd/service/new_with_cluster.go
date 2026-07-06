package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store"
)

func NewWithCluster(
	s *store.Store,
	c *cluster.Cluster,
) *Service {
	return &Service{
		clusters: map[string]*cluster.Cluster{c.Name(): c},
		store:    s,
	}
}
