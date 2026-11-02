package watcher

import (
	chat "github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/socket"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/peer/hold"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"time"
)

func New(
	c chat.MattermostClient,
	g *peer.Guard,
	s *store.Store,
	l *logger.Logger,
	r face.Reporter,
) *Watcher {
	result := &Watcher{
		client:   c,
		guard:    g,
		store:    s,
		logger:   l,
		reporter: r,
		index:    map[string]string{},
		buffer:   map[string][]*event.Event{},
		since:    map[string]time.Time{},
	}
	result.hold = hold.New(c.Host(), socket.New(c, result.Dispatch), l, r)

	return result
}
