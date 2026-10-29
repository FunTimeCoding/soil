//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/generative/ollama/reachable_skip"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"testing"
)

func openTestStore(t *testing.T) (*store.Store, *ollama.Client) {
	t.Helper()
	o := ollama.NewEnvironment()
	t.Cleanup(o.Close)
	reachable_skip.Skip(t)

	return store_tester.OpenTestStore(t), o
}
