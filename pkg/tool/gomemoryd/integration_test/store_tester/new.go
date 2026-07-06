package store_tester

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := store.New(filepath.Join(t.TempDir(), constant.TestDatabase))
	t.Cleanup(s.Close)

	return &Tester{t: t, Store: s}
}
