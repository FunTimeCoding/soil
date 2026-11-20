package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"testing"
)

func TestResolveAbsolutePath(t *testing.T) {
	s := store_tester.New(t).IndexedTestStore()
	path := s.ResolveAbsolutePath("test", "alfa.md")
	assert.StringContains(t, "alfa.md", path)
	assert.StringContains(t, "/", path)
}
