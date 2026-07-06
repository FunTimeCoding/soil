package service_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/mock_notifier"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	memoryMock "github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_indexer"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := store_tester.New(t)
	c := mock_client.New()
	si := mock_indexer.New()
	ci := mock_indexer.New()
	n := mock_notifier.New()
	harbor := t.TempDir()

	return &Tester{
		Store: s,
		t:     t,
		Service: service.New(
			s.Store,
			c,
			memoryMock.New(),
			si,
			ci,
			n,
			memory.New(),
			harbor,
			s.Clock(),
			logger.New(context.Background()),
		),
		Harbor:            harbor,
		Client:            c,
		SummaryIndexer:    si,
		CompletionIndexer: ci,
		Notifier:          n,
	}
}
