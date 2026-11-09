package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/pointer/retarget"
	"github.com/funtimecoding/soil/pkg/lint/pointer_tester"
	"github.com/funtimecoding/soil/pkg/lint/types/symbol_move"
	"testing"
)

func TestRetargetRenamedTestCarriesBareCitation(t *testing.T) {
	after, changes := lint.RetargetDocument(
		pointer_tester.Resolver(),
		"doc/note.md",
		"Pinned by `TestShelfHoldsStock` and `TestShelfHoldsStock()`, not `TestShelfHoldsStockLong`; the TestShelfHoldsStock prose stays.\n",
		symbolMoves(
			symbol_move.New(
				"pkg/shop/unit",
				"TestShelfHoldsStock",
				"pkg/shop/unit",
				"TestShelfKeepsStock",
			),
		),
	)
	assert.String(
		t,
		"Pinned by `TestShelfKeepsStock` and `TestShelfKeepsStock()`, not `TestShelfHoldsStockLong`; the TestShelfHoldsStock prose stays.\n",
		after,
	)
	assert.Count(t, 2, changes)
}

func TestRetargetRepeatedTestNameIsLeftForHand(t *testing.T) {
	content := "Pinned by `TestGuard`.\n"
	r := pointer_tester.Resolver()
	r.TestElsewhere = func(
		name string,
		directory string,
	) bool {
		return name == "TestGuard" && directory == "pkg/shop/unit"
	}
	after, changes := lint.RetargetDocument(
		r,
		"doc/note.md",
		content,
		symbolMoves(
			symbol_move.New(
				"pkg/shop/unit",
				"TestGuard",
				"pkg/shop/unit",
				"TestShopGuard",
			),
		),
	)
	assert.String(t, content, after)
	assert.Count(t, 1, changes)
	assert.String(
		t,
		"test name is declared in another package too",
		changes[0].Reason,
	)
}

func TestRetargetMentionsDocumentCitingRenamedTest(t *testing.T) {
	moves := symbolMoves(
		symbol_move.New(
			"pkg/shop/unit",
			"TestShelfHoldsStock",
			"pkg/shop/unit",
			"TestShelfKeepsStock",
		),
	)
	assert.True(
		t,
		retarget.Mentions(
			"doc/note.md",
			"Pinned by `TestShelfHoldsStock`.\n",
			moves,
		),
	)
	assert.False(
		t,
		retarget.Mentions(
			"doc/note.md",
			"Pinned by `TestShelfCounts`.\n",
			moves,
		),
	)
}
