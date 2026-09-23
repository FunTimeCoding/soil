package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"testing"
)

func NewStore(t *testing.T) *store.Store {
	t.Helper()
	result := store.New(lite.NewMemory())
	t.Cleanup(result.Close)

	return result
}
