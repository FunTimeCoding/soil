package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFindSimilarFiles(t *testing.T) {
	s := indexedTestStore(t)
	defer s.Close()
	similar := s.MustFindSimilarFiles("test/alfa.md", 5)
	assert.Greater(t, 0, float64(len(similar)))
	assert.String(t, "test/alpha.md", similar[0])
}

func TestFindSimilarFilesNoMatch(t *testing.T) {
	s := indexedTestStore(t)
	defer s.Close()
	similar := s.MustFindSimilarFiles("completely-unrelated-path", 5)
	assert.Count(t, 0, similar)
}
