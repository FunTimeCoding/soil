package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestResolveAbsolutePath(t *testing.T) {
	s := indexedTestStore(t)
	defer s.Close()
	path := s.ResolveAbsolutePath("test", "alpha.md")
	assert.StringContains(t, "alpha.md", path)
	assert.StringContains(t, "/", path)
}
