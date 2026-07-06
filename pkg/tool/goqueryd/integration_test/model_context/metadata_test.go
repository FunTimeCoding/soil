//go:build local

package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration_test/model_context_tester"
	"testing"
)

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
