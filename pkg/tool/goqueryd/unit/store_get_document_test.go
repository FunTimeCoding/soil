package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"testing"
)

func TestGetDocumentByRelativePath(t *testing.T) {
	s := store_tester.IndexedTestStore(t)
	defer s.Close()
	d := s.MustGetDocument("test/alfa.md")
	assert.NotNil(t, d)
	assert.String(t, "Search Pipeline", d.Title)
	assert.StringContains(t, "hybrid search pipeline", d.Body)
}

func TestGetDocumentByVirtualPath(t *testing.T) {
	s := store_tester.IndexedTestStore(t)
	defer s.Close()
	d := s.MustGetDocument("qmd://test/alfa.md")
	assert.NotNil(t, d)
	assert.String(t, "Search Pipeline", d.Title)
	assert.String(t, "qmd://test/alfa.md", d.VirtualPath)
}

func TestFindDocumentMissIsNotFound(t *testing.T) {
	s := store_tester.IndexedTestStore(t)
	defer s.Close()
	d, found, e := s.FindDocument("test/nonexistent.md")
	assert.FatalOnError(t, e)
	assert.False(t, found)
	assert.Nil(t, d)
}

func TestGetDocumentWithContext(t *testing.T) {
	s := store_tester.IndexedTestStore(t)
	defer s.Close()
	s.AddContext("test", constant.Slash, "root context")
	d := s.MustGetDocument("test/alfa.md")
	assert.NotNil(t, d)
	assert.String(t, "root context", d.Context)
}
