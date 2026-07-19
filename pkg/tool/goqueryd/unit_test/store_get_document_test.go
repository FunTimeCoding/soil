package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"testing"
)

func TestGetDocumentByRelativePath(t *testing.T) {
	s := indexedTestStore(t)
	defer s.Close()
	d := s.MustGetDocument("test/alpha.md")
	assert.NotNil(t, d)
	assert.String(t, "Search Pipeline", d.Title)
	assert.StringContains(t, "hybrid search pipeline", d.Body)
}

func TestGetDocumentByVirtualPath(t *testing.T) {
	s := indexedTestStore(t)
	defer s.Close()
	d := s.MustGetDocument("qmd://test/alpha.md")
	assert.NotNil(t, d)
	assert.String(t, "Search Pipeline", d.Title)
	assert.String(t, "qmd://test/alpha.md", d.VirtualPath)
}

func TestGetDocumentNotFound(t *testing.T) {
	s := indexedTestStore(t)
	defer s.Close()
	d := s.MustGetDocument("test/nonexistent.md")
	assert.Nil(t, d)
}

func TestGetDocumentWithContext(t *testing.T) {
	s := indexedTestStore(t)
	defer s.Close()
	s.AddContext("test", separator.Slash, "root context")
	d := s.MustGetDocument("test/alpha.md")
	assert.NotNil(t, d)
	assert.String(t, "root context", d.Context)
}
