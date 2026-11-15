package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/highlight"
	"github.com/funtimecoding/soil/pkg/web/web_tester"
	"testing"
)

func TestHighlightMarksEveryTermIgnoringCase(t *testing.T) {
	assert.String(
		t,
		`B<mark class="highlight">oar</mark>s an <mark class="highlight">Oar</mark>`,
		web_tester.Render(t, highlight.New("Boars an Oar", "oar").Node()),
	)
}

func TestHighlightPrefersLongerTermAtSameStart(t *testing.T) {
	assert.String(
		t,
		`<mark class="highlight">lifting</mark> on`,
		web_tester.Render(
			t,
			highlight.New("lifting on", "lift", "lifting").Node(),
		),
	)
}

func TestHighlightWithoutTermsLeavesTextPlain(t *testing.T) {
	assert.String(t, "alfa", web_tester.Render(t, highlight.New("alfa").Node()))
}

func TestHighlightSkipsEmptyTerm(t *testing.T) {
	assert.String(
		t,
		`<mark class="highlight">alfa</mark> bravo`,
		web_tester.Render(t, highlight.New("alfa bravo", "", "alfa").Node()),
	)
}

func TestMarkedHighlightFollowsMarksInText(t *testing.T) {
	assert.String(
		t,
		`…crews <mark class="highlight">rowing</mark> home, <mark class="highlight">rows</mark> of boats`,
		web_tester.Render(
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
		web_tester.Render(
			t,
			highlight.NewMarked("alfa [bravo", "[", "]").Node(),
		),
	)
}

func TestHighlightLeavesTextWhoseLowerCaseChangesLength(t *testing.T) {
	assert.String(
		t,
		"İstanbul",
		web_tester.Render(t, highlight.New("İstanbul", "stan").Node()),
	)
}
