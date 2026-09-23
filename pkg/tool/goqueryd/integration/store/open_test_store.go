//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/generative/ollama/reachable_skip"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"testing"
)

func openTestStore(t *testing.T) (*store.Store, *ollama.Client) {
	t.Helper()
	o := ollama.NewEnvironment()
	reachable_skip.Skip(t)
	result := store.New(connection.NewMemory())
	t.Cleanup(result.Close)

	return result, o
}
