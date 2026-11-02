package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	prometheus "github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"time"
)

func New(
	s *store.Store,
	a prometheus.AlertmanagerClient,
	g *peer.Guard,
	retention time.Duration,
	n face.EventNotifier,
) *Service {
	return &Service{
		store:        s,
		alertmanager: a,
		guard:        g,
		retention:    retention,
		notifier:     n,
	}
}
