package service_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/mock_notifier"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	memoryMock "github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_indexer"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := store_tester.New(t)
	c := mock_client.New()
	mc := memoryMock.New()
	si := mock_indexer.New()
	ci := mock_indexer.New()
	n := mock_notifier.New()
	x := search_index.New(connection.NewMemory())
	r := memory.New()
	harbor := t.TempDir()
	c.Harbor = harbor

	return &Tester{
		Store: s,
		t:     t,
		Service: service.New(
			s.Store,
			c,
			mc,
			si,
			ci,
			x,
			n,
			r,
			harbor,
			s.Clock(),
			logger.New(context.Background()),
		),
		Harbor:            harbor,
		Client:            c,
		SummaryIndexer:    si,
		CompletionIndexer: ci,
		Notifier:          n,
		Memory:            mc,
		Search:            x,
		Reporter:          r,
	}
}
