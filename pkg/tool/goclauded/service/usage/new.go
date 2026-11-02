package usage

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	library "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"time"
)

func New(
	s *store.Store,
	n goclauded.Notifier,
	browser *chromium.Client,
	clock func() time.Time,
	l *logger.Logger,
	r library.Reporter,
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
