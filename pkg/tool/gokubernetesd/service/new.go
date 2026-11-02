package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/port_forward/registry"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store"
)

func New(
	clusters map[string]*cluster.Cluster,
	guards map[string]*peer.Guard,
	s *store.Store,
	forwards *registry.Registry,
	r face.Reporter,
) *Service {
	return &Service{
		clusters: clusters,
		guards:   guards,
		store:    s,
		forwards: forwards,
		reporter: r,
	}
}
