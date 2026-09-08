package sweeper

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"time"
)

func New(
	s *store.Store,
	n *notifier.Notifier,
	f face.Forgetter,
	l *logger.Logger,
	r soilFace.Reporter,
	window time.Duration,
	interval time.Duration,
) *Sweeper {
	return &Sweeper{
		store:     s,
		notifier:  n,
		forgetter: f,
		logger:    l,
		reporter:  r,
		window:    window,
		interval:  interval,
	}
}
