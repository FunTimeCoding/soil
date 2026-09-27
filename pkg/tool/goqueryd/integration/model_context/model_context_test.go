//go:build local

package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration/model_context_tester"
	"strings"
	"testing"
)

func TestAddAndListContexts(t *testing.T) {
	s := model_context_tester.New(t)
	s.AddContext("docs", "/design/", "Architecture and design decisions")
	result := s.ListContexts()
	assert.StringContains(t, "design", result)
	assert.StringContains(t, "Architecture", result)
}

func TestRemoveContext(t *testing.T) {
	s := model_context_tester.New(t)
	s.AddContext("docs", "/temp/", "Temporary context")
	s.RemoveContext("docs", "/temp/")
	result := s.ListContexts()
	assert.StringNotContains(t, "Temporary", result)
}

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

func TestGetDocument(t *testing.T) {
	s := indexFixtures(t)
	result := s.Get("test/alfa.md")
	assert.StringContains(t, "Search Pipeline", result)
	assert.StringContains(t, "hybrid search pipeline", result)
}

func TestGetDocumentVirtualPath(t *testing.T) {
	s := indexFixtures(t)
	result := s.Get("qmd://test/alfa.md")
	assert.StringContains(t, "Search Pipeline", result)
}

func TestGetDocumentNotFoundSuggestsAlternatives(t *testing.T) {
	s := indexFixtures(t)
	result := s.MustCallToolError(
		constant.Get,
		map[string]any{constant.Path: "test/alfaa.md"},
	)
	assert.StringContains(t, "Did you mean", result)
}

func TestAddCollectionAndIndex(t *testing.T) {
	s := model_context_tester.New(t)
	s.IndexFixtures()
	result := s.MustCallTool(constant.Index, map[string]any{})
	assert.StringContains(t, "test", result)
}

func TestStatusAfterIndex(t *testing.T) {
	s := model_context_tester.New(t)
	s.IndexFixtures()
	result := s.MustCallTool(constant.Status, map[string]any{})
	assert.StringContains(t, "total_documents", result)
}

func TestListDocuments(t *testing.T) {
	s := indexFixtures(t)
	result := s.ListDocuments("test")
	assert.StringContains(t, "alfa.md", result)
	assert.StringContains(t, "tools/charlie.md", result)
	assert.StringContains(t, "archive/echo.md", result)
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

func TestPushMetadataAppearsInKeywordSearch(t *testing.T) {
	s := model_context_tester.New(t)
	s.PushWithMetadata(
		"notes",
		"flavored.md",
		"# Flavored\n\nA document with vanilla metadata.\n",
		map[string]string{"flavor": "vanilla"},
	)
	result := s.SearchKeyword("vanilla metadata")
	assert.StringContains(t, "Flavored", result)
	assert.StringContains(t, "vanilla", result)
}

func TestMetadataFilterIncludes(t *testing.T) {
	s := model_context_tester.New(t)
	s.PushWithMetadata(
		"notes",
		"apple.md",
		"# Apple\n\nA fruit document.\n",
		map[string]string{"kind": "fruit"},
	)
	s.PushWithMetadata(
		"notes",
		"carrot.md",
		"# Carrot\n\nA vegetable document.\n",
		map[string]string{"kind": "vegetable"},
	)
	result := s.SearchKeywordWithMetadata(
		"document",
		map[string]string{"kind": "fruit"},
	)
	assert.StringContains(t, "Apple", result)
	assert.StringNotContains(t, "Carrot", result)
}

func TestMetadataFilterExcludes(t *testing.T) {
	s := model_context_tester.New(t)
	s.PushWithMetadata(
		"notes",
		"only.md",
		"# Only\n\nThe only document here.\n",
		map[string]string{"color": "red"},
	)
	result := s.SearchKeywordWithMetadata(
		"only document",
		map[string]string{"color": "blue"},
	)
	assert.StringNotContains(t, "Only", result)
}

func TestSourceTypeViaMetadata(t *testing.T) {
	s := model_context_tester.New(t)
	s.PushWithMetadata(
		"notes",
		"typed.md",
		"# Typed\n\nSource type set via metadata.\n",
		map[string]string{constant.SourceType: "custom-type"},
	)
	result := s.MustCallTool(
		constant.Search,
		map[string]any{
			"query":             "source type metadata",
			constant.Mode:       "keyword",
			constant.SourceType: "custom-type",
		},
	)
	assert.StringContains(t, "Typed", result)
}

func TestDeleteCascadesMetadata(t *testing.T) {
	s := model_context_tester.New(t)
	s.PushWithMetadata(
		"notes",
		"doomed.md",
		"# Doomed\n\nWill be deleted with metadata.\n",
		map[string]string{"ephemeral": "true"},
	)
	result := s.SearchKeywordWithMetadata(
		"deleted metadata",
		map[string]string{"ephemeral": "true"},
	)
	assert.StringContains(t, "Doomed", result)
	s.Delete("notes", "doomed.md")
	after := s.SearchKeywordWithMetadata(
		"deleted metadata",
		map[string]string{"ephemeral": "true"},
	)
	assert.StringNotContains(t, "Doomed", after)
}

func TestDeleteCollectionCascadesMetadata(t *testing.T) {
	s := model_context_tester.New(t)
	s.PushWithMetadata(
		"ephemeral",
		"temp.md",
		"# Temporary\n\nCollection will be deleted.\n",
		map[string]string{"lifespan": "short"},
	)
	result := s.SearchKeywordWithMetadata(
		"collection deleted",
		map[string]string{"lifespan": "short"},
	)
	assert.StringContains(t, "Temporary", result)
	s.MustCallTool(
		constant.DeleteCollection,
		map[string]any{"name": "ephemeral"},
	)
	after := s.SearchKeywordWithMetadata(
		"collection deleted",
		map[string]string{"lifespan": "short"},
	)
	assert.StringNotContains(t, "Temporary", after)
}

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

func TestSearchKeyword(t *testing.T) {
	s := indexFixtures(t)
	result := s.SearchKeyword("hybrid search pipeline")
	assert.StringContains(t, "Search Pipeline", result)
}

func TestSearchKeywordCollectionFilter(t *testing.T) {
	s := indexFixtures(t)
	result := s.MustCallTool(
		constant.Search,
		map[string]any{
			"query":             "chunking",
			constant.Mode:       "keyword",
			constant.Collection: "test",
		},
	)
	assert.StringContains(t, "Chunking Strategy", result)
}

func TestTagSetAndGet(t *testing.T) {
	s := model_context_tester.New(t)
	s.SetTag("design/", "design-doc")
	result := s.GetTag("design/")
	assert.StringContains(t, "design-doc", result)
}

func TestTagRemove(t *testing.T) {
	s := model_context_tester.New(t)
	s.SetTag("temp/", "scratch")
	s.SetTag("temp/", "")
	result := s.GetTag("temp/")
	assert.StringContains(t, "no tag", result)
}
