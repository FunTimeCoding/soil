package monitor

import (
	mattermost "github.com/funtimecoding/soil/pkg/chat/mattermost/face"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor/scheduler"
	"github.com/mattermost/mattermost/server/public/model"
)

func New(
	c mattermost.Client,
	g *peer.Guard,
	o *option.Monitor,
	l *logger.Logger,
	r face.Reporter,
) *Monitor {
	m := &Monitor{
		client:               c,
		guard:                g,
		configuration:        o,
		logger:               l,
		reporter:             r,
		channelCache:         make(map[string]*model.Channel),
		lastCheckMillisecond: make(map[string]int64),
	}
	m.scheduler = scheduler.New(o.Schedule, m.Poll, l, r)

	return m
}
