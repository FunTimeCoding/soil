//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestResolveAbsolutePath(t *testing.T) {
	s, _ := indexedTestStore(t)
	defer s.Close()
	path := s.ResolveAbsolutePath("test", "alpha.md")
	assert.StringContains(t, "alpha.md", path)
	assert.StringContains(t, "/", path)
}
