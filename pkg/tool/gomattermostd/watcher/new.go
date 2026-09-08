package watcher

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	mattermostFace "github.com/funtimecoding/soil/pkg/tool/gomattermostd/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"time"
)

func New(
	c mattermostFace.ChatSource,
	s *store.Store,
	n *notifier.Notifier,
	l *logger.Logger,
	r face.Reporter,
	window time.Duration,
) *Watcher {
	return &Watcher{
		client:   c,
		store:    s,
		notifier: n,
		logger:   l,
		reporter: r,
		window:   window,
		index:    map[string]string{},
		buffer:   map[string][]*event.Event{},
		timer:    map[string]*time.Timer{},
	}
}
