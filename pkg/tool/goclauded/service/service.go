package service

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	library "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/session_cache"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"sync"
	"sync/atomic"
	"time"
)

type Service struct {
	store             *store.Store
	client            face.ClaudeSource
	memory            client.Client
	summaryIndexer    queryd.Indexer
	completionIndexer queryd.Indexer
	search            *search_index.Index
	searchIndexed     atomic.Int64
	searchTotal       atomic.Int64
	notifier          face.Notifier
	reporter          library.Reporter
	clock             func() time.Time
	logger            *logger.Logger
	harbor            string
	cache             *session_cache.Cache
	browser           *chromium.Client
	browserMutex      sync.Mutex
	lastMemoryPoll    string
	done              chan struct{}
	doneOnce          sync.Once
}
