package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/worker"
	"gorm.io/gorm"
	"sync"
)

type Daemon struct {
	database  *gorm.DB
	service   *service.Service
	web       *web.Server
	watcher   *worker.Watcher
	closeOnce sync.Once
}
