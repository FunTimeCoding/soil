//go:build local

package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration_test/service_tester"
	"testing"
)

func TestGetDocument(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	document, similar, e := s.Service.GetDocument("test/alpha.md")
	assert.FatalOnError(t, e)
	assert.NotNil(t, document)
	assert.String(t, "Search Pipeline", document.Title)
	assert.Count(t, 0, similar)
}

func TestGetDocumentNotFoundWithSuggestions(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	document, similar, e := s.Service.GetDocument("test/alfa.md")
	assert.FatalOnError(t, e)
	assert.Nil(t, document)
	assert.Greater(t, 0, float64(len(similar)))
}

func TestGetDocumentNotFoundNoSuggestions(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	document, similar, e := s.Service.GetDocument("completely-unrelated")
	assert.FatalOnError(t, e)
	assert.Nil(t, document)
	assert.Count(t, 0, similar)
}
