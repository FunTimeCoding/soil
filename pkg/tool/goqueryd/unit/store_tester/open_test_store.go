package store_tester

import (
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"testing"
)

func OpenTestStore(t *testing.T) *store.Store {
	t.Helper()

	return store.New(connection.NewMemory())
}
