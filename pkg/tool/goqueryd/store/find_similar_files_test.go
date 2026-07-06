//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFindSimilarFiles(t *testing.T) {
	s, _ := indexedTestStore(t)
	defer s.Close()
	similar := s.MustFindSimilarFiles("test/alfa.md", 5)
	assert.Greater(t, 0, float64(len(similar)))
	assert.String(t, "test/alpha.md", similar[0])
}

func TestFindSimilarFilesNoMatch(t *testing.T) {
	s, _ := indexedTestStore(t)
	defer s.Close()
	similar := s.MustFindSimilarFiles("completely-unrelated-path", 5)
	assert.Count(t, 0, similar)
}
