package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/document"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/event"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/label"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/memory"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/roster"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/search"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/startup"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/transcript"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/usage"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/session_cache"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stopper"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/double"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	directory := t.TempDir()
	l := logger.New(t.Context())
	s := store_tester.NewOnFile(t, filepath.Join(directory, "goclauded.sqlite"))
	d := double.New(t)
	x := connection.New(l, filepath.Join(directory, constant.SearchIndexFile))
	result := &Tester{
		Set:            d,
		Store:          s,
		t:              t,
		Index:          search_index.New(x),
		searchDatabase: x,
		memoryGuard: peer.New(
			constant.MemoryPeer,
			d.MemoryPeer.Address(),
			d.MemoryPeer,
			l,
			d.Reporter,
		),
		queryGuard: peer.New(
			constant.QueryPeer,
			d.Query.Address(),
			d.Query,
			l,
			d.Reporter,
		),
		stop: stopper.New(),
	}
	t.Cleanup(result.Close)
	result.Label = label.New(s.Store, d.Notifier)
	result.Message = message.New(
		s.Store,
		d.Notifier,
		result.stop,
		s.Clock(),
		d.Reporter,
	)
	result.Memory = memory.New(
		s.Store,
		d.MemoryPeer,
		result.memoryGuard,
		result.Message,
		s.Clock(),
		d.Reporter,
	)
	result.Usage = usage.New(s.Store, d.Notifier, nil, s.Clock(), l, d.Reporter)
	result.Search = search.New(s.Store, result.Index, d.Harbor)
	result.Document = document.New(s.Store, d.Query, result.queryGuard, l)
	result.Event = event.New(
		s.Store,
		result.Message,
		result.Memory,
		result.Document,
		d.Notifier,
		s.Clock(),
		d.Reporter,
	)
	sources := func() []string { return []string{d.Source} }
	cache := session_cache.New()
	result.Transcript = transcript.New(
		s.Store,
		d.Client,
		cache,
		result.Memory,
		result.Index,
		d.Harbor,
		sources,
		t.TempDir(),
		l,
		d.Reporter,
	)
	result.Session = session.New(
		s.Store,
		result.Label,
		result.Message,
		result.Transcript,
		d.Client,
		cache,
		result.Index,
		d.Notifier,
		s.Clock(),
		d.Harbor,
		sources,
	)
	result.Roster = roster.New(
		s.Store,
		result.Message,
		result.Session,
		cache,
		d.Notifier,
		s.Clock(),
		l,
	)
	result.Startup = startup.New(
		result.Roster,
		result.Transcript,
		result.Search,
	)

	return result
}
