package watcher

import (
	mattermostFace "github.com/funtimecoding/soil/pkg/chat/mattermost/face"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/peer/reacher"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"time"
)

func New(
	c mattermostFace.Client,
	g *peer.Guard,
	s *store.Store,
	n *notifier.Notifier,
	l *logger.Logger,
	r face.Reporter,
	window time.Duration,
) *Watcher {
	return &Watcher{
		client:   c,
		guard:    g,
		store:    s,
		notifier: n,
		logger:   l,
		reporter: r,
		reacher:  reacher.New(r),
		host:     c.Host(),
		window:   window,
		index:    map[string]string{},
		buffer:   map[string][]*event.Event{},
		timer:    map[string]*time.Timer{},
	}
}
