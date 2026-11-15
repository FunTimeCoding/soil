package store_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"testing"
)

func Open(t *testing.T) (*store.Store, *ollama.Client) {
	t.Helper()
	o := client(t)

	return store_tester.OpenTestStore(t), o
}
