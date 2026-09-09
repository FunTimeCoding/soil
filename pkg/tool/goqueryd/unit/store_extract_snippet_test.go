package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestExtractSnippetFindsQueryTerms(t *testing.T) {
	snippet, line := store.ExtractSnippet(
		constant.TestBody,
		"keyword matching",
		0,
	)
	assert.StringContains(t, "Keyword Matching", snippet)
	assert.StringContains(t, "term frequency", snippet)
	assert.Greater(t, 0, line)
}

func TestExtractSnippetWithChunkPosition(t *testing.T) {
	position := strings.Index(constant.TestBody, "## Cross-Encoder Reranking")
	snippet, line := store.ExtractSnippet(
		constant.TestBody,
		"cross-encoder scores pair",
		position,
	)
	assert.StringContains(t, "Cross-Encoder Reranking", snippet)
	assert.Greater(t, 0, line)
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
	assert.Less(t, 401, len(snippet))
}

func TestExtractSnippetTruncatesOnRuneBoundary(t *testing.T) {
	long := strings.Join(
		[]string{
			"keyword",
			strings.Repeat("a", 387),
			strings.Repeat("—", 20),
		},
		" ",
	)
	snippet, _ := store.ExtractSnippet(long, "keyword", 0)
	assert.True(t, utf8.ValidString(snippet))
	assert.Less(t, 401, len(snippet))
}

func TestExtractSnippetChunkWindowOnRuneBoundary(t *testing.T) {
	body := strings.Join(
		[]string{strings.Repeat("—", 100), "keyword target line"},
		"",
	)
	position := strings.Index(body, "keyword")
	snippet, _ := store.ExtractSnippet(body, "keyword", position)
	assert.True(t, utf8.ValidString(snippet))
	assert.StringContains(t, "keyword", snippet)
}
