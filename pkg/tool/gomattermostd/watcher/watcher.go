package watcher

import (
	mattermostFace "github.com/funtimecoding/soil/pkg/chat/mattermost/face"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/peer/hold"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"sync"
	"time"
)

type Watcher struct {
	client   mattermostFace.Client
	guard    *peer.Guard
	store    *store.Store
	logger   *logger.Logger
	reporter face.Reporter
	hold     *hold.Hold
	self     string
	index    map[string]string
	buffer   map[string][]*event.Event
	since    map[string]time.Time
	running  bool
	once     sync.Once
	mutex    sync.Mutex
}
