package service

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/face"
	nextcloud "github.com/funtimecoding/soil/pkg/nextcloud/face"
	"github.com/funtimecoding/soil/pkg/peer"
	metric "github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/store"
	"sync"
)

type Service struct {
	board           *board.Board
	store           *store.Store
	prometheus      metric.PrometheusClient
	prometheusGuard *peer.Guard
	argocd          kubernetes.ArgocdClient
	argocdGuard     *peer.Guard
	usage           nextcloud.UsageClient
	usageGuard      *peer.Guard
	notifier        *notifier.Notifier
	labels          map[string]bool
	mutex           sync.RWMutex
	values          map[string][]string
}
