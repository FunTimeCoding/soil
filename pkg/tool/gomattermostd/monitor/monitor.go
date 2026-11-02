package monitor

import (
	"github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor/option"
	"github.com/mattermost/mattermost/server/public/model"
	"sync"
)

type Monitor struct {
	client               face.MattermostClient
	guard                *peer.Guard
	configuration        *option.Monitor
	logger               *logger.Logger
	channelCache         map[string]*model.Channel
	lastCheckMillisecond map[string]int64
	notifyChannel        *model.Channel
	username             string
	initialized          bool
	mutex                sync.RWMutex
}
