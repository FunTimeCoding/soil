package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"testing"
)

func TestListDocuments(t *testing.T) {
	s := store_tester.IndexedTestStore(t)
	defer s.Close()
	entries := s.MustListDocuments("test")
	assert.Count(t, 5, entries)
}

func TestListDocumentsEmptyCollection(t *testing.T) {
	s := store_tester.OpenTestStore(t)
	defer s.Close()
	entries := s.MustListDocuments("nonexistent")
	assert.Count(t, 0, entries)
}
