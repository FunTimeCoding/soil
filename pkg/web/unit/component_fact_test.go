package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/fact"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"testing"
)

func TestFactTableLabelsEachRow(t *testing.T) {
	assert.String(
		t,
		`<table class="facts"><tbody><tr><th scope="row">Alfa</th><td>bravo</td></tr></tbody></table>`,
		renderedFact(t, fact.New().Row("Alfa", "bravo")),
	)
}

func TestSectionSpansBothColumns(t *testing.T) {
	assert.String(
		t,
		`<table class="facts"><tbody><tr class="section"><th colspan="2">Charlie</th></tr><tr><th scope="row">Alfa</th><td>bravo</td></tr></tbody></table>`,
		renderedFact(t, fact.New().Section("Charlie").Row("Alfa", "bravo")),
	)
}

func TestFactTooltipLandsOnLastValue(t *testing.T) {
	assert.String(
		t,
		`<table class="facts"><tbody><tr><th scope="row">Alfa</th><td>bravo</td></tr><tr><th scope="row">Charlie</th><td><span data-tooltip="echo">delta</span></td></tr></tbody></table>`,
		renderedFact(
			t,
			fact.New().
				Row("Alfa", "bravo").
				Row("Charlie", "delta").
				Tooltip(constant.FixtureEcho),
		),
	)
}

func TestGlanceSetsLabelOverValue(t *testing.T) {
	assert.String(
		t,
		`<div class="glance"><div><div class="glance-label">Alfa</div><div class="glance-value">bravo</div></div><div><div class="glance-label">Charlie</div><div class="glance-value" data-tooltip="echo">delta</div></div></div>`,
		renderedFact(
			t,
			fact.New().
				Row("Alfa", "bravo").
				Row("Charlie", "delta").
				Tooltip(constant.FixtureEcho).
				Glance(),
		),
	)
}

func TestGlanceColumnsCarryCount(t *testing.T) {
	assert.String(
		t,
		`<div class="glance glance-columns" style="--glance-columns: 3"><div><div class="glance-label">Alfa</div><div class="glance-value">bravo</div></div></div>`,
		renderedFact(t, fact.New().Row("Alfa", "bravo").Columns(3).Glance()),
	)
}

func TestGlanceSectionsBecomeGroups(t *testing.T) {
	assert.String(
		t,
		`<div><div class="glance-group"><div class="glance"><div><div class="glance-label">Alfa</div><div class="glance-value">bravo</div></div></div></div><div class="glance-group"><h5>Charlie</h5><div class="glance"><div><div class="glance-label">Delta</div><div class="glance-value">echo</div></div></div></div></div>`,
		renderedFact(
			t,
			fact.New().
				Row("Alfa", "bravo").
				Section("Charlie").
				Row("Delta", constant.FixtureEcho).
				Glance(),
		),
	)
}

func TestFactValueCanBeMarkup(t *testing.T) {
	assert.StringContains(
		t,
		`<tr><th scope="row">Alfa</th><td><a href="/delta">delta</a></td></tr>`,
		renderedFact(
			t,
			fact.New().RowNode(
				"Alfa",
				html.A(html.Href("/delta"), gomponents.Text("delta")),
			),
		),
	)
}
