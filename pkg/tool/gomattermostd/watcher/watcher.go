package watcher

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	mattermostFace "github.com/funtimecoding/soil/pkg/tool/gomattermostd/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"sync"
	"time"
)

type Watcher struct {
	client   mattermostFace.ChatSource
	store    *store.Store
	notifier *notifier.Notifier
	logger   *logger.Logger
	reporter face.Reporter
	window   time.Duration
	self     string
	index    map[string]string
	buffer   map[string][]*event.Event
	timer    map[string]*time.Timer
	done     chan struct{}
	running  bool
	once     sync.Once
	mutex    sync.Mutex
}
