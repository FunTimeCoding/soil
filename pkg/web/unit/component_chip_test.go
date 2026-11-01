package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"testing"
)

func TestAChipIsASpanALinkOrAButton(t *testing.T) {
	assert.String(
		t,
		`<span class="chip">alfa</span>`,
		renderedNode(t, chip.New("alfa").Node()),
	)
	assert.String(
		t,
		`<a href="/bravo" class="chip service" style="--chip-color: #3d9970">bravo<small>healthy</small></a>`,
		renderedNode(
			t,
			chip.New("bravo").
				Link("/bravo").
				Kind("service").
				Color("#3d9970").
				Content(html.Small(gomponents.Text("healthy"))).
				Node(),
		),
	)
}

func TestAToggleChipIsAPressableButtonCarryingItsHooks(t *testing.T) {
	assert.String(
		t,
		`<button type="button" aria-pressed="true" class="chip toggle active" title="charlie help" data-tag="charlie" onclick="toggleTag(this)">charlie</button>`,
		renderedNode(
			t,
			chip.New("charlie").
				Toggle(true).
				Title("charlie help").
				Dataset(constant.FixtureTag, "charlie").
				OnClick("toggleTag(this)").
				Node(),
		),
	)
	assert.StringContains(
		t,
		`aria-pressed="false" class="chip toggle"`,
		renderedNode(t, chip.New("delta").Toggle(false).Node()),
	)
}

func TestACurrentLinkChipCarriesTheActiveLookAndTheMark(t *testing.T) {
	assert.String(
		t,
		`<a href="/golf" class="chip active" aria-current="true">golf</a>`,
		renderedNode(t, chip.New("golf").Link("/golf").Current().Node()),
	)
}
