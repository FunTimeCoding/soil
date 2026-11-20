package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/badge"
	"github.com/funtimecoding/soil/pkg/web/web_tester"
	"maragu.dev/gomponents"
	"testing"
)

func TestBadgeIsPillInPrimaryColour(t *testing.T) {
	assert.String(
		t,
		`<span class="badge">alfa</span>`,
		web_tester.New(t).Render(badge.New("alfa").Node()),
	)
}

func TestQuietBadgeAndColouredOneCarryTheirState(t *testing.T) {
	assert.String(
		t,
		`<span class="badge quiet">bravo</span>`,
		web_tester.New(t).Render(badge.New("bravo").Quiet().Node()),
	)
	assert.String(
		t,
		`<span class="badge" style="--badge-color: #6e40c9">charlie</span>`,
		web_tester.New(t).Render(badge.New("charlie").Color("#6e40c9").Node()),
	)
}

func TestBadgeRowWrapsItsBadgesAndVanishesWhenEmpty(t *testing.T) {
	assert.String(
		t,
		`<div class="badge-row"><span class="badge">delta</span></div>`,
		web_tester.New(t).Render(badge.Row(badge.New("delta").Node())),
	)
	assert.String(
		t,
		"",
		web_tester.New(t).Render(badge.Row([]gomponents.Node{}...)),
	)
}
