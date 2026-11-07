package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/lease"
	atlas "github.com/funtimecoding/soil/pkg/tool/goatlasd/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/service"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/worker"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"gorm.io/gorm"
	"sync"
)

type Daemon struct {
	database    *gorm.DB
	netbox      netbox.Client
	netboxGuard *peer.Guard
	peers       []face.Closer
	guards      []*peer.Guard
	collectors  []atlas.Collector
	lease       *lease.Collector
	service     *service.Service
	web         *web.Server
	clock       *worker.Worker
	logger      *logger.Logger
	reporter    face.Reporter
	closeOnce   sync.Once
}
