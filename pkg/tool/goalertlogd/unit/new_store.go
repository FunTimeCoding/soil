package unit

import (
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"testing"
)

func newStore(t *testing.T) *store.Store {
	t.Helper()
	d := lite.NewMemory()
	t.Cleanup(func() { relational.CloseMapper(d) })

	return store.New(d)
}
