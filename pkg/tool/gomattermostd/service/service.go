package service

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/face"
	"github.com/funtimecoding/soil/pkg/peer"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/sweeper"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
	"gorm.io/gorm"
	"sync"
)

type Service struct {
	database   *gorm.DB
	store      *store.Store
	mattermost face.Client
	guard      *peer.Guard
	connector  goclauded.Client
	watcher    *watcher.Watcher
	sweeper    *sweeper.Sweeper
	monitor    *monitor.Monitor
	closeOnce  sync.Once
}
