package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stopper"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/mock_notifier"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	memoryMock "github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	queryMock "github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	directory := t.TempDir()
	l := logger.New(t.Context())
	s := store_tester.NewOnFile(t, filepath.Join(directory, "goclauded.sqlite"))
	c := mock_client.New()
	mc := memoryMock.New()
	q := queryMock.New()
	n := mock_notifier.New()
	r := memory.New()
	harbor := t.TempDir()
	source := t.TempDir()
	c.Harbor = harbor
	x := connection.New(l, filepath.Join(directory, constant.SearchIndexFile))
	result := &Tester{
		Store:          s,
		t:              t,
		Harbor:         harbor,
		Source:         source,
		Client:        c,
		Query:          q,
		Notifier:       n,
		Memory:         mc,
		Search:         search_index.New(x),
		Reporter:       r,
		searchDatabase: x,
		memoryGuard:    peer.New(constant.MemoryPeer, mc.Address(), mc, l, r),
		queryGuard:     peer.New(constant.QueryPeer, q.Address(), q, l, r),
		stop:           stopper.New(),
	}
	t.Cleanup(result.Close)
	result.Service = service.New(
		s.Store,
		result.Search,
		c,
		mc,
		result.memoryGuard,
		q,
		result.queryGuard,
		n,
		nil,
		result.stop,
		s.Clock(),
		t.TempDir(),
		harbor,
		func() []string { return []string{source} },
		false,
		l,
		r,
	)

	return result
}
