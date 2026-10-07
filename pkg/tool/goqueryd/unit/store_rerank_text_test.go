package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_reranker"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"testing"
)

func TestLineOffset(t *testing.T) {
	assert.Integer(t, 0, store.LineOffset("a\nbb\nccc", 0))
	assert.Integer(t, 0, store.LineOffset("a\nbb\nccc", 1))
	assert.Integer(t, 5, store.LineOffset("a\nbb\nccc", 3))
	assert.Integer(t, 5, store.LineOffset("a\nbb\nccc", 9))
}

func TestRerankTextKeywordHitTakesTheChunkHoldingItsLine(t *testing.T) {
	assert.StringNotContains(
		t,
		"bravo",
		store.RerankText(
			halves(),
			result.NewSearch("halves.md", "", constant.NoChunkPosition, 1),
			mock_reranker.New(),
		),
	)
	assert.StringContains(
		t,
		"bravo",
		store.RerankText(
			halves(),
			result.NewSearch("halves.md", "", constant.NoChunkPosition, 1200),
			mock_reranker.New(),
		),
	)
}

func TestRerankTextVectorChunkWinsOverSnippetLine(t *testing.T) {
	assert.StringNotContains(
		t,
		"bravo",
		store.RerankText(
			halves(),
			result.NewSearch("halves.md", "", 0, 1200),
			mock_reranker.New(),
		),
	)
}

func TestRerankTextWithoutBodyUsesSnippet(t *testing.T) {
	assert.String(
		t,
		"charlie",
		store.RerankText(
			"",
			result.NewSearch(
				"halves.md",
				"charlie",
				constant.NoChunkPosition,
				0,
			),
			mock_reranker.New(),
		),
	)
}
