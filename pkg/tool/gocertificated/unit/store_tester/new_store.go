package store_tester

import (
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
	"testing"
)

func NewStore(t *testing.T) *store.Store {
	t.Helper()
	d := lite.NewMemory()
	t.Cleanup(func() { relational.CloseMapper(d) })

	return store.New(d)
}
