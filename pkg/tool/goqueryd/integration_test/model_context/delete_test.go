//go:build local

package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration_test/model_context_tester"
	"testing"
)

func TestDeleteDocument(t *testing.T) {
	s := model_context_tester.New(t)
	s.Push("notes", "doomed.md", "# Doomed\n\nWill be deleted.\n")
	s.Delete("notes", "doomed.md")
	result := s.MustCallToolError(
		constant.Get,
		map[string]any{constant.Path: "notes/doomed.md"},
	)
	assert.StringContains(t, "not found", result)
}

func TestDeleteDocumentNotFound(t *testing.T) {
	s := model_context_tester.New(t)
	result := s.MustCallToolError(
		constant.Delete,
		map[string]any{
			constant.Collection: "notes",
			constant.Path:       "nonexistent.md",
		},
	)
	assert.StringContains(t, "not found", result)
}
