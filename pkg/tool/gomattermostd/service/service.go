package service

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/face"
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Service struct {
	database    *gorm.DB
	store       *store.Store
	mattermost  face.Client
	guard       *peer.Guard
	connector   goclauded.Client
	rosterGuard *peer.Guard
	notifier    *notifier.Notifier
	watcher     *watcher.Watcher
	monitor     *monitor.Monitor
	logger      *logger.Logger
	reporter    soilFace.Reporter
	debounce    time.Duration
	purge       time.Duration
	closeOnce   sync.Once
}
