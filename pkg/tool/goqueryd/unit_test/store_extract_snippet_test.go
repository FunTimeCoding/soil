package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"strings"
	"testing"
)

func TestExtractSnippetFindsQueryTerms(t *testing.T) {
	snippet, line := store.ExtractSnippet(
		constant.TestBody,
		"keyword matching",
		0,
	)
	assert.StringContains(t, "Keyword Matching", snippet)
	assert.StringContains(t, "term frequency", snippet)
	assert.Greater(t, 0, float64(line))
}

func TestExtractSnippetWithChunkPosition(t *testing.T) {
	position := strings.Index(constant.TestBody, "## Cross-Encoder Reranking")
	snippet, line := store.ExtractSnippet(
		constant.TestBody,
		"cross-encoder scores pair",
		position,
	)
	assert.StringContains(t, "Cross-Encoder Reranking", snippet)
	assert.Greater(t, 0, float64(line))
}

func TestExtractSnippetEmptyBody(t *testing.T) {
	snippet, line := store.ExtractSnippet("", "anything", 0)
	assert.String(t, "", snippet)
	assert.Float(t, 0, float64(line))
}

func TestExtractSnippetTruncatesLongResult(t *testing.T) {
	long := strings.Repeat(
		"keyword matching vector similarity cross-encoder reranking\n",
		50,
	)
	snippet, _ := store.ExtractSnippet(long, "keyword", 0)
	assert.Less(
		t,
		float64(constant.SnippetMaxLength+1),
		float64(len(snippet)),
	)
}
