package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/badge"
	"maragu.dev/gomponents"
	"testing"
)

func TestABadgeIsAPillInThePrimaryColour(t *testing.T) {
	assert.String(
		t,
		`<span class="badge">alfa</span>`,
		renderedNode(t, badge.New("alfa").Node()),
	)
}

func TestAQuietBadgeAndAColouredOneCarryTheirState(t *testing.T) {
	assert.String(
		t,
		`<span class="badge quiet">bravo</span>`,
		renderedNode(t, badge.New("bravo").Quiet().Node()),
	)
	assert.String(
		t,
		`<span class="badge" style="--badge-color: #6e40c9">charlie</span>`,
		renderedNode(t, badge.New("charlie").Color("#6e40c9").Node()),
	)
}

func TestABadgeRowWrapsItsBadgesAndVanishesWhenEmpty(t *testing.T) {
	assert.String(
		t,
		`<div class="badge-row"><span class="badge">delta</span></div>`,
		renderedNode(t, badge.Row(badge.New("delta").Node())),
	)
	assert.String(t, "", renderedNode(t, badge.Row([]gomponents.Node{}...)))
}
