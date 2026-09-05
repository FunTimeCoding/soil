//go:build local

package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"testing"
)

func TestGetDocument(t *testing.T) {
	s := indexFixtures(t)
	result := s.Get("test/alpha.md")
	assert.StringContains(t, "Search Pipeline", result)
	assert.StringContains(t, "hybrid search pipeline", result)
}

func TestGetDocumentVirtualPath(t *testing.T) {
	s := indexFixtures(t)
	result := s.Get("qmd://test/alpha.md")
	assert.StringContains(t, "Search Pipeline", result)
}

func TestGetDocumentNotFoundSuggestsAlternatives(t *testing.T) {
	s := indexFixtures(t)
	result := s.MustCallToolError(
		constant.Get,
		map[string]any{constant.Path: "test/alfa.md"},
	)
	assert.StringContains(t, "Did you mean", result)
}
