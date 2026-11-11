package usage

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"time"
)

func New(
	s *store.Store,
	n face.EventNotifier,
	browser *chromium.Client,
	clock func() time.Time,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	return &Service{
		store:    s,
		notifier: n,
		browser:  browser,
		clock:    clock,
		logger:   l,
		reporter: r,
	}
}
