package daemon

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/runner"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/service"
	"gorm.io/gorm"
	"sync"
)

type Daemon struct {
	database   *gorm.DB
	reaper     *reaper.Reaper
	downstream []face.Client
	guards     []*peer.Guard
	runner     *runner.Runner
	service    *service.Service
	closeOnce  sync.Once
}
