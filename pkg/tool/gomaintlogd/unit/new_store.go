package unit

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store"
	"testing"
)

func newStore(
	t *testing.T,
	n *notifier.Notifier,
) *store.Store {
	t.Helper()
	result := store.New(lite.NewMemory(), n)
	t.Cleanup(result.Close)

	return result
}
