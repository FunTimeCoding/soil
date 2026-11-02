package service

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	library "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/session_cache"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stopper"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	memoryd "github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"time"
)

func New(
	s *store.Store,
	search *search_index.Index,
	c face.ClaudeSource,
	m memoryd.Client,
	memoryGuard *peer.Guard,
	q queryd.Client,
	queryGuard *peer.Guard,
	n face.Notifier,
	browser *chromium.Client,
	stop *stopper.Stopper,
	clock func() time.Time,
	exportPath string,
	harbor string,
	sources func() []string,
	utilizationFallback bool,
	l *logger.Logger,
	r library.Reporter,
) *Service {
	return &Service{
		exportPath:          exportPath,
		store:               s,
		client:              c,
		memory:              m,
		memoryGuard:         memoryGuard,
		query:               q,
		queryGuard:          queryGuard,
		search:              search,
		notifier:            n,
		reporter:            r,
		harbor:              harbor,
		sources:             sources,
		cache:              session_cache.New(),
		clock:               clock,
		logger:              l,
		browser:             browser,
		utilizationFallback: utilizationFallback,
		stop:                stop,
		lastMemoryPoll:      clock().UTC().Format("2006-01-02T15:04:05Z"),
	}
}
