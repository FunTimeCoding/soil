package service

import (
	chat "github.com/funtimecoding/soil/pkg/chat/face"
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
	"time"
)

type Service struct {
	store      *store.Store
	mattermost chat.MattermostClient
	guard      *peer.Guard
	notifier   *notifier.Notifier
	watcher    *watcher.Watcher
	monitor    *monitor.Monitor
	logger     *logger.Logger
	reporter   soilFace.Reporter
	debounce   time.Duration
	purge      time.Duration
}
