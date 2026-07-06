package watcher

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
)

func New(
	s *service.Service,
	l *logger.Logger,
	r face.Reporter,
	seedDirectory string,
) *Watcher {
	return &Watcher{
		service:       s,
		recovery:      recovery.New(l, r),
		seedDirectory: seedDirectory,
		stop:          make(chan struct{}),
	}
}
