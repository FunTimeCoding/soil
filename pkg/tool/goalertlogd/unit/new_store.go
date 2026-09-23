package unit

import (
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"testing"
)

func newStore(t *testing.T) *store.Store {
	t.Helper()
	result := store.New(lite.NewMemory())
	t.Cleanup(result.Close)

	return result
}
