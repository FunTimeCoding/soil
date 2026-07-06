package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store"
	"sync"
)

type Service struct {
	clusters     map[string]*cluster.Cluster
	store        *store.Store
	portForwards sync.Map
}
