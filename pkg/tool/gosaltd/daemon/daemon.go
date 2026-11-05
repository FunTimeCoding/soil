package daemon

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/provision/salt/face"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/runner"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/service"
	"gorm.io/gorm"
	"sync"
)

type Daemon struct {
	database  *gorm.DB
	reaper    *reaper.Reaper
	salt      face.SaltClient
	guard     *peer.Guard
	runner    *runner.Runner
	service   *service.Service
	closeOnce sync.Once
}
