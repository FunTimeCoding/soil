package service

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/face"
	nextcloud "github.com/funtimecoding/soil/pkg/nextcloud/face"
	"github.com/funtimecoding/soil/pkg/peer"
	metric "github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/store"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"gorm.io/gorm"
	"sync"
)

type Service struct {
	board           *board.Board
	database        *gorm.DB
	store           *store.Store
	prometheus      metric.Client
	prometheusGuard *peer.Guard
	argocd          kubernetes.ArgocdClient
	argocdGuard     *peer.Guard
	usage           nextcloud.UsageClient
	usageGuard      *peer.Guard
	authorization   *client.Client
	notifier        *notifier.Notifier
	labels          map[string]bool
	mutex           sync.RWMutex
	values          map[string][]string
	closeOnce       sync.Once
}
