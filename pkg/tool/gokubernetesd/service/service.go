package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/port_forward"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store"
	"gorm.io/gorm"
	"sync"
)

type Service struct {
	database     *gorm.DB
	clusters     map[string]*cluster.Cluster
	guards       map[string]*peer.Guard
	store        *store.Store
	reporter     face.Reporter
	mutex        sync.Mutex
	portForwards map[string]*port_forward.Forward
}
