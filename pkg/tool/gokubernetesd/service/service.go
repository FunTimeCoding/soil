package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/port_forward/registry"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store"
)

type Service struct {
	clusters map[string]*cluster.Cluster
	guards   map[string]*peer.Guard
	store    *store.Store
	forwards *registry.Registry
	reporter face.Reporter
}
