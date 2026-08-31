package service

import (
	library "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/session_cache"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"time"
)

type Service struct {
	store             *store.Store
	client            face.ClaudeSource
	memory            client.Client
	summaryIndexer    queryd.Indexer
	completionIndexer queryd.Indexer
	notifier          face.Notifier
	reporter          library.Reporter
	clock             func() time.Time
	logger            *logger.Logger
	harbor            string
	cache             *session_cache.Cache
	lastMemoryPoll    string
}
