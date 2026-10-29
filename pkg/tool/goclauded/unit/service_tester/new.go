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
	queryMock "github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := store_tester.New(t)
	c := mock_client.New()
	mc := memoryMock.New()
	q := queryMock.New()
	t.Cleanup(q.Close)
	n := mock_notifier.New()
	x := search_index.New(connection.NewMemory())
	r := memory.New()
	harbor := t.TempDir()
	c.Harbor = harbor
	result := &Tester{
		Store: s,
		t:     t,
		Service: service.New(
			s.Store,
			c,
			mc,
			q,
			x,
			n,
			r,
			harbor,
			s.Clock(),
			logger.New(context.Background()),
		),
		Harbor:   harbor,
		Client:   c,
		Query:    q,
		Notifier: n,
		Memory:   mc,
		Search:   x,
		Reporter: r,
	}
	t.Cleanup(result.Service.Close)

	return result
}
