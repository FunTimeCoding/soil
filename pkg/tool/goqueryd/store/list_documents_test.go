//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestListDocuments(t *testing.T) {
	s, _ := indexedTestStore(t)
	defer s.Close()
	entries := s.MustListDocuments("test")
	assert.Count(t, 5, entries)
}

func TestListDocumentsEmptyCollection(t *testing.T) {
	s, _ := openTestStore(t)
	defer s.Close()
	entries := s.MustListDocuments("nonexistent")
	assert.Count(t, 0, entries)
}
