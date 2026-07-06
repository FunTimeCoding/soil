//go:build local

package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration_test/model_context_tester"
	"testing"
)

func TestPushAndSearch(t *testing.T) {
	s := model_context_tester.New(t)
	s.Push("notes", "first.md", "# First Note\n\nPushed directly via MCP.\n")
	result := s.SearchKeyword("pushed directly")
	assert.StringContains(t, "First Note", result)
}

func TestPushAndGet(t *testing.T) {
	s := model_context_tester.New(t)
	s.Push(
		"notes",
		"retrievable.md",
		"# Retrievable\n\nContent that can be fetched back.\n",
	)
	result := s.Get("notes/retrievable.md")
	assert.StringContains(t, "Retrievable", result)
	assert.StringContains(t, "fetched back", result)
}

func TestPushWithSourceType(t *testing.T) {
	s := model_context_tester.New(t)
	s.MustCallTool(
		constant.Push,
		map[string]any{
			constant.Collection: "memories",
			constant.Path:       "memory/1",
			constant.Body:       "# Test Memory\n\nA pushed memory.\n",
			constant.SourceType: "memory",
		},
	)
	result := s.MustCallTool(
		constant.Search,
		map[string]any{
			"query":             "pushed memory",
			constant.Mode:       "keyword",
			constant.SourceType: "memory",
		},
	)
	assert.StringContains(t, "Test Memory", result)
}
