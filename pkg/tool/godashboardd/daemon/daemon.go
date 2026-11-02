package daemon

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	nextcloud "github.com/funtimecoding/soil/pkg/nextcloud/face"
	"github.com/funtimecoding/soil/pkg/peer"
	metric "github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/service"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/web"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"gorm.io/gorm"
	"sync"
)

type Daemon struct {
	database        *gorm.DB
	prometheus      metric.PrometheusClient
	prometheusGuard *peer.Guard
	argocd          kubernetes.ArgocdClient
	argocdGuard     *peer.Guard
	usage           nextcloud.UsageClient
	usageGuard      *peer.Guard
	authorization   *client.Client
	notifier        *notifier.Notifier
	service         *service.Service
	web             *web.Server
	logger          *logger.Logger
	reporter        face.Reporter
	closeOnce       sync.Once
}
