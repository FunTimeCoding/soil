package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"testing"
)

func TestResolveAbsolutePath(t *testing.T) {
	s := store_tester.IndexedTestStore(t)
	defer s.Close()
	path := s.ResolveAbsolutePath("test", "alpha.md")
	assert.StringContains(t, "alpha.md", path)
	assert.StringContains(t, "/", path)
}
