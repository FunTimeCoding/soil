package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"testing"
)

func TestLineOffset(t *testing.T) {
	assert.Integer(t, 0, store.LineOffset("a\nbb\nccc", 0))
	assert.Integer(t, 0, store.LineOffset("a\nbb\nccc", 1))
	assert.Integer(t, 5, store.LineOffset("a\nbb\nccc", 3))
	assert.Integer(t, 5, store.LineOffset("a\nbb\nccc", 9))
}

func TestRerankTextKeywordHitTakesChunkHoldingItsLine(t *testing.T) {
	assert.StringNotContains(
		t,
		"bravo",
		store.RerankText(
			store_tester.Halves(),
			result.NewSearch("halves.md", "", constant.NoChunkPosition, 1),
			service_tester.NewReranker(t),
		),
	)
	assert.StringContains(
		t,
		"bravo",
		store.RerankText(
			store_tester.Halves(),
			result.NewSearch("halves.md", "", constant.NoChunkPosition, 1200),
			service_tester.NewReranker(t),
		),
	)
}

func TestRerankTextVectorChunkWinsOverSnippetLine(t *testing.T) {
	assert.StringNotContains(
		t,
		"bravo",
		store.RerankText(
			store_tester.Halves(),
			result.NewSearch("halves.md", "", 0, 1200),
			service_tester.NewReranker(t),
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
			service_tester.NewReranker(t),
		),
	)
}
