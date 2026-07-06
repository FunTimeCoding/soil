package scheduler

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/robfig/cron/v3"
	"sync"
)

type Scheduler struct {
	cron     *cron.Cron
	schedule string
	task     func()
	logger   *logger.Logger
	reporter face.Reporter
	entry    cron.EntryID
	mutex    sync.Mutex
	running  bool
}
