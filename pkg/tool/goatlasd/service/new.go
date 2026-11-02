package service

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/lease"
	atlas "github.com/funtimecoding/soil/pkg/tool/goatlasd/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"time"
)

func New(
	s *store.Store,
	n netbox.Client,
	netboxGuard *peer.Guard,
	collectors []atlas.Collector,
	l *lease.Collector,
	retention time.Duration,
) *Service {
	return &Service{
		store:       s,
		netbox:      n,
		netboxGuard: netboxGuard,
		collectors:  collectors,
		lease:       l,
		retention:   retention,
	}
}
