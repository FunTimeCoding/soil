package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"testing"
)

func TestFindSimilarFiles(t *testing.T) {
	s := store_tester.IndexedTestStore(t)
	defer s.Close()
	similar := s.MustFindSimilarFiles("test/alfa.md", 5)
	assert.Greater(t, 0, len(similar))
	assert.String(t, "test/alpha.md", similar[0])
}

func TestFindSimilarFilesNoMatch(t *testing.T) {
	s := store_tester.IndexedTestStore(t)
	defer s.Close()
	similar := s.MustFindSimilarFiles("completely-unrelated-path", 5)
	assert.Count(t, 0, similar)
}
