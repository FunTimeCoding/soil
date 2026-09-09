package monitor

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	mattermost "github.com/funtimecoding/soil/pkg/tool/gomattermostd/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor/scheduler"
	"github.com/mattermost/mattermost/server/public/model"
	"sync"
)

type Monitor struct {
	client               mattermost.MattermostSource
	configuration        *option.Monitor
	logger               *logger.Logger
	reporter             face.Reporter
	scheduler            *scheduler.Scheduler
	channelCache         map[string]*model.Channel
	lastCheckMillisecond map[string]int64
	notifyChannel        *model.Channel
	username             string
	running              bool
	mutex                sync.RWMutex
}
