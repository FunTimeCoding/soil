package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/highlight"
	"testing"
)

func TestHighlightMarksEveryTermIgnoringCase(t *testing.T) {
	assert.String(
		t,
		`B<mark class="highlight">oar</mark>s an <mark class="highlight">Oar</mark>`,
		renderedNode(t, highlight.New("Boars an Oar", "oar").Node()),
	)
}

func TestHighlightPrefersLongerTermAtSameStart(t *testing.T) {
	assert.String(
		t,
		`<mark class="highlight">lifting</mark> on`,
		renderedNode(t, highlight.New("lifting on", "lift", "lifting").Node()),
	)
}

func TestHighlightWithoutTermsLeavesTextPlain(t *testing.T) {
	assert.String(t, "alfa", renderedNode(t, highlight.New("alfa").Node()))
}

func TestHighlightSkipsEmptyTerm(t *testing.T) {
	assert.String(
		t,
		`<mark class="highlight">alfa</mark> bravo`,
		renderedNode(t, highlight.New("alfa bravo", "", "alfa").Node()),
	)
}

func TestMarkedHighlightFollowsMarksInText(t *testing.T) {
	assert.String(
		t,
		`…crews <mark class="highlight">rowing</mark> home, <mark class="highlight">rows</mark> of boats`,
		renderedNode(
			t,
			highlight.NewMarked(
				"…crews [rowing] home, [rows] of boats",
				"[",
				"]",
			).
				Node(),
		),
	)
}

func TestMarkedHighlightDropsUnclosedMark(t *testing.T) {
	assert.String(
		t,
		"alfa bravo",
		renderedNode(t, highlight.NewMarked("alfa [bravo", "[", "]").Node()),
	)
}

func TestHighlightLeavesTextWhoseLowerCaseChangesLength(t *testing.T) {
	assert.String(
		t,
		"İstanbul",
		renderedNode(t, highlight.New("İstanbul", "stan").Node()),
	)
}
