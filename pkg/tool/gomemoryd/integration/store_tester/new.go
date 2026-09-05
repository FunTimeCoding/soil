package store_tester

import (
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := store.New(connection.NewMemory())
	t.Cleanup(s.Close)

	return &Tester{t: t, Store: s}
}
