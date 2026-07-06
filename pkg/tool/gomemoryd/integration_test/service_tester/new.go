package service_tester

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_indexer"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := store.New(filepath.Join(t.TempDir(), constant.TestDatabase))
	t.Cleanup(s.Close)
	i := mock_indexer.New()

	return &Tester{Service: service.New(s, i, i, i), Indexer: i}
}
