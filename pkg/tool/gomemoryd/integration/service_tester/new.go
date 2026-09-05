package service_tester

import (
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_indexer"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := store.New(connection.NewMemory())
	t.Cleanup(s.Close)
	i := mock_indexer.New()

	return &Tester{Service: service.New(s, i, i, i), Indexer: i}
}
