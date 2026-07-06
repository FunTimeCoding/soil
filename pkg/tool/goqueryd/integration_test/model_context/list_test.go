//go:build local

package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration_test/model_context_tester"
	"strings"
	"testing"
)

func TestListDocuments(t *testing.T) {
	s := indexFixtures(t)
	result := s.ListDocuments("test")
	assert.StringContains(t, "alpha.md", result)
	assert.StringContains(t, "tools/gamma.md", result)
	assert.StringContains(t, "archive/epsilon.md", result)
}

func TestListBySourceType(t *testing.T) {
	s := model_context_tester.New(t)
	s.PushWithMetadata(
		"notes",
		"summary.md",
		"# Summary\n\nA session summary.\n",
		map[string]string{"source_type": "session-summary"},
	)
	s.PushWithMetadata(
		"notes",
		"completion.md",
		"# Completion\n\nA session completion.\n",
		map[string]string{"source_type": "session-completion"},
	)
	filtered := s.ListWithFilter("notes", "session-summary", 0, true)
	assert.StringContains(t, "summary.md", filtered)
	assert.StringNotContains(t, "completion.md", filtered)
	all := s.ListDocuments("notes")
	assert.StringContains(t, "summary.md", all)
	assert.StringContains(t, "completion.md", all)
}

func TestListWithLimit(t *testing.T) {
	s := model_context_tester.New(t)
	s.Push("notes", "a.md", "# A\n\nFirst.\n")
	s.Push("notes", "b.md", "# B\n\nSecond.\n")
	s.Push("notes", "c.md", "# C\n\nThird.\n")
	result := s.ListWithFilter("notes", "", 2, false)
	assert.Integer(t, 2, strings.Count(result, "qmd://notes/"))
}
