package watcher

import (
	mattermostFace "github.com/funtimecoding/soil/pkg/chat/mattermost/face"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/socket"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/peer/hold"
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
	result := &Watcher{
		client:   c,
		guard:    g,
		store:    s,
		notifier: n,
		logger:   l,
		reporter: r,
		window:   window,
		index:    map[string]string{},
		buffer:   map[string][]*event.Event{},
		timer:    map[string]*time.Timer{},
	}
	result.hold = hold.New(c.Host(), socket.New(c, result.Dispatch), l, r)

	return result
}
