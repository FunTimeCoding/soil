package store_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/telegram/store"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"testing"
)

func New(t *testing.T) *store.Store {
	t.Helper()
	result := store.New(lite.NewMemory())
	t.Cleanup(result.Close)

	return result
}
