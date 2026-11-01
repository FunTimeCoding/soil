package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/option"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/session_cache"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stopper"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"path/filepath"
)

func New(
	o *option.Option,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	d := lite.New(l, o.LitePath)
	x := connection.New(
		l,
		filepath.Join(filepath.Dir(o.LitePath), constant.SearchIndexFile),
	)
	m := o.Memory()
	q := o.Query()

	return &Service{
		database:       d,
		searchDatabase: x,
		exportPath:     o.SessionExportPath,
		store:          store.New(d, o.Clock),
		client:         o.Claude(),
		memory:         m,
		memoryGuard:    peer.New(constant.MemoryPeer, m.Address(), m, l, r),
		query:          q,
		queryGuard:     peer.New(constant.QueryPeer, q.Address(), q, l, r),
		search:         search_index.New(x),
		notifier:       o.Notifier(),
		reporter:       r,
		harbor:         o.Harbor,
		cache:          session_cache.New(),
		clock:          o.Clock,
		logger:         l,
		stop:           stopper.New(),
		lastMemoryPoll: o.Clock().UTC().Format("2006-01-02T15:04:05Z"),
	}
}
