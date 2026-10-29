package service_tester

import (
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := store.New(connection.NewMemory())
	t.Cleanup(s.Close)
	q := mock_client.New()
	t.Cleanup(q.Close)

	return &Tester{t: t, Service: service.New(s, q), Indexer: q}
}
