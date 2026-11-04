package daemon

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/port_forward/registry"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"gorm.io/gorm"
	"sync"
)

type Daemon struct {
	database  *gorm.DB
	clusters  map[string]*cluster.Cluster
	guards    map[string]*peer.Guard
	forwards  *registry.Registry
	service   *service.Service
	armed     bool
	closeOnce sync.Once
}
