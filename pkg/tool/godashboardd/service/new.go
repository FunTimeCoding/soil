package service

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/face"
	nextcloud "github.com/funtimecoding/soil/pkg/nextcloud/face"
	"github.com/funtimecoding/soil/pkg/peer"
	metric "github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/store"
)

func New(
	b *board.Board,
	s *store.Store,
	p metric.PrometheusClient,
	prometheusGuard *peer.Guard,
	a kubernetes.ArgocdClient,
	argocdGuard *peer.Guard,
	u nextcloud.UsageClient,
	usageGuard *peer.Guard,
	n *notifier.Notifier,
) *Service {
	result := &Service{
		board:           b,
		store:           s,
		prometheus:      p,
		prometheusGuard: prometheusGuard,
		argocd:          a,
		argocdGuard:     argocdGuard,
		usage:           u,
		usageGuard:      usageGuard,
		notifier:        n,
		labels:          map[string]bool{},
		values:          map[string][]string{},
	}

	for _, entry := range b.Entries() {
		result.labels[entry.Label] = true
	}

	return result
}
