package monitor

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor/option"
	"github.com/mattermost/mattermost/server/public/model"
)

func New(
	c face.Client,
	g *peer.Guard,
	o *option.Monitor,
	l *logger.Logger,
) *Monitor {
	return &Monitor{
		client:               c,
		guard:                g,
		configuration:        o,
		logger:               l,
		channelCache:         make(map[string]*model.Channel),
		lastCheckMillisecond: make(map[string]int64),
	}
}
