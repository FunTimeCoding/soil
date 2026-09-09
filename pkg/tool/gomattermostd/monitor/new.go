package monitor

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	mattermost "github.com/funtimecoding/soil/pkg/tool/gomattermostd/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor/scheduler"
	"github.com/mattermost/mattermost/server/public/model"
)

func New(
	c mattermost.MattermostSource,
	o *option.Monitor,
	l *logger.Logger,
	r face.Reporter,
) *Monitor {
	m := &Monitor{
		client:               c,
		configuration:        o,
		logger:               l,
		reporter:             r,
		channelCache:         make(map[string]*model.Channel),
		lastCheckMillisecond: make(map[string]int64),
	}
	m.scheduler = scheduler.New(o.Schedule, m.run, l, r)

	return m
}
