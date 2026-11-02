package daemon

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/runner"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/service"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"gorm.io/gorm"
	"sync"
)

type Daemon struct {
	database       *gorm.DB
	reaper         *reaper.Reaper
	runner         *runner.Runner
	inventory      face.Client
	inventoryGuard *peer.Guard
	service        *service.Service
	closeOnce      sync.Once
}
