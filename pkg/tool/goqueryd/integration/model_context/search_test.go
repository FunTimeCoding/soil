//go:build local

package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"testing"
)

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
