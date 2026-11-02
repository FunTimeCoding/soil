package daemon

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/gauge"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stopper"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/watcher"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web"
	memoryd "github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"gorm.io/gorm"
	"sync"
)

type Daemon struct {
	database       *gorm.DB
	searchDatabase *sql.DB
	memory         memoryd.Client
	memoryGuard    *peer.Guard
	query          queryd.Client
	queryGuard     *peer.Guard
	browser        *chromium.Client
	stop           *stopper.Stopper
	service        *service.Service
	web            *web.Server
	gauge          *gauge.Gauge
	watcher        *watcher.Watcher
	clocks         []face.Worker
	logger         *logger.Logger
	closeOnce      sync.Once
}
