package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"testing"
)

func TestATableHeadsItsColumnsAndCarriesItsRows(t *testing.T) {
	assert.String(
		t,
		`<table><thead><tr><th>Alfa</th><th class="numeric">Bravo</th></tr></thead><tbody><tr><td>charlie</td><td class="numeric">1</td></tr></tbody></table>`,
		renderedTable(
			t,
			table.New("Alfa", "Bravo").
				Numeric(1).
				Rows(
					[]gomponents.Node{
						html.Tr(
							html.Td(gomponents.Text("charlie")),
							html.Td(
								html.Class(constant.NumericClass),
								gomponents.Text("1"),
							),
						),
					},
				).
				Empty("Nothing yet."),
		),
	)
}

func TestAnEmptyTableSaysSoInTheSharedLook(t *testing.T) {
	assert.String(
		t,
		`<p class="absent">Nothing yet.</p>`,
		renderedTable(t, table.New("Alfa").Empty("Nothing yet.")),
	)
}

func TestATableWithoutAnEmptyMessageStaysATable(t *testing.T) {
	assert.String(
		t,
		`<table><thead><tr><th>Alfa</th></tr></thead><tbody></tbody></table>`,
		renderedTable(t, table.New("Alfa")),
	)
}

func TestAHeaderlessTableCarriesOnlyItsBody(t *testing.T) {
	assert.String(
		t,
		`<table class="echo"><tbody><tr><td>alfa</td></tr></tbody></table>`,
		renderedTable(
			t,
			table.New().
				Headerless().
				Class(constant.FixtureEcho).
				Rows(
					[]gomponents.Node{html.Tr(html.Td(gomponents.Text("alfa")))},
				),
		),
	)
}

func TestAnIdentifiedTableKeepsItsTargetWhenEmpty(t *testing.T) {
	assert.String(
		t,
		`<div id="delta"><p class="absent">Nothing yet.</p></div>`,
		renderedTable(
			t,
			table.New("Alfa").Empty("Nothing yet.").Identifier("delta"),
		),
	)
}

func TestATableCarriesItsClassAndBodyIdentifier(t *testing.T) {
	assert.String(
		t,
		`<table class="echo"><thead><tr><th>Alfa</th></tr></thead><tbody id="foxtrot"></tbody></table>`,
		renderedTable(
			t,
			table.New("Alfa").Class(constant.FixtureEcho).BodyIdentifier(
				"foxtrot",
			),
		),
	)
}
