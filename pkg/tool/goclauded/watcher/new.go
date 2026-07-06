package watcher

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
)

func New(
	s *service.Service,
	l *logger.Logger,
	r face.Reporter,
	harbor string,
) *Watcher {
	return &Watcher{
		service:  s,
		recovery: recovery.New(l, r),
		logger:   l,
		harbor:   harbor,
		changed:  make(map[string]string),
	}
}
