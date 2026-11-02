package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/types/discovered_file"
)

func New(
	sync func([]discovered_file.File),
	seedDirectory string,
	l *logger.Logger,
	r face.Reporter,
) *Watcher {
	return &Watcher{
		sync:          sync,
		recovery:      recovery.New(l, r),
		reporter:      r,
		seedDirectory: seedDirectory,
		stop:          make(chan struct{}),
	}
}
