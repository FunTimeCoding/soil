package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/goflightd/service"
	"github.com/funtimecoding/soil/pkg/worker"
	"gorm.io/gorm"
	"sync"
)

type Daemon struct {
	database  *gorm.DB
	service   *service.Service
	clocks    []*worker.Worker
	closeOnce sync.Once
}
