package scheduler

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/robfig/cron/v3"
	"sync"
)

type Scheduler struct {
	cron     *cron.Cron
	schedule string
	task     func()
	recovery *recovery.Recovery
	entry    cron.EntryID
	mutex    sync.Mutex
	running  bool
}
