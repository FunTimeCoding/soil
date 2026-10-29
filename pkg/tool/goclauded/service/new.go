package service

import (
	library "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/session_cache"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"time"
)

func New(
	s *store.Store,
	c face.ClaudeSource,
	m client.Client,
	q queryd.Client,
	x *search_index.Index,
	n face.Notifier,
	r library.Reporter,
	harbor string,
	clock func() time.Time,
	l *logger.Logger,
) *Service {
	return &Service{
		store:             s,
		client:            c,
		memory:            m,
		memoryGuard:       peer.New(constant.MemoryPeer, m.Address(), m, l, r),
		query:             q,
		search:            x,
		notifier:          n,
		reporter:          r,
		harbor:            harbor,
		cache:             session_cache.New(),
		clock:             clock,
		logger:            l,
		done:              make(chan struct{}),
		lastMemoryPoll:    clock().UTC().Format("2006-01-02T15:04:05Z"),
	}
}
