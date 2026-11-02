package daemon

import (
	chat "github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/peer"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
	"github.com/funtimecoding/soil/pkg/worker"
	"gorm.io/gorm"
	"sync"
)

type Daemon struct {
	database    *gorm.DB
	mattermost  chat.MattermostClient
	guard       *peer.Guard
	connector   goclauded.Client
	rosterGuard *peer.Guard
	watcher     *watcher.Watcher
	service     *service.Service
	clocks      []*worker.Worker
	closeOnce   sync.Once
}
