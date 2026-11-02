package daemon

import (
	directory "github.com/funtimecoding/soil/pkg/directory/face"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/tool/gogated/web"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"gorm.io/gorm"
	"sync"
)

type Daemon struct {
	database       *gorm.DB
	directory      directory.DirectoryClient
	directoryGuard *peer.Guard
	authorization  *client.Client
	service        *service.Service
	web            *web.Server
	logger         *logger.Logger
	reporter       face.Reporter
	closeOnce      sync.Once
}
