//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/system/constant"
	constant1 "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"testing"
)

func indexedTestStore(t *testing.T) (*store.Store, *ollama.Client) {
	t.Helper()
	s, o := openTestStore(t)
	s.AddCollection(
		"test",
		fixture.Path(constant.SearchPath),
		constant1.DefaultGlob,
	)
	s.Index("test")

	return s, o
}
