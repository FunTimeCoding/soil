package daemon

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/web"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
	"gorm.io/gorm"
	"sync"
)

type Daemon struct {
	database     *gorm.DB
	alertmanager face.AlertmanagerClient
	guard        *peer.Guard
	service      *service.Service
	web          *web.Server
	clock        *worker.Worker
	closeOnce    sync.Once
}
