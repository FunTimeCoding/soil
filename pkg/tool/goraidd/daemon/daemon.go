package daemon

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/service"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/web"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/worker"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"gorm.io/gorm"
	"sync"
)

type Daemon struct {
	database      *gorm.DB
	parser        face.Client
	parserGuard   *peer.Guard
	authorization *client.Client
	service       *service.Service
	web           *web.Server
	watcher       *worker.Worker
	logger        *logger.Logger
	reporter      soilFace.Reporter
	closeOnce     sync.Once
}
