package service

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/lease"
	atlas "github.com/funtimecoding/soil/pkg/tool/goatlasd/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"time"
)

type Service struct {
	store       *store.Store
	netbox      netbox.Client
	netboxGuard *peer.Guard
	collectors  []atlas.Collector
	lease       *lease.Collector
	retention   time.Duration
}
