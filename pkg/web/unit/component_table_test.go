package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/web_tester"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"testing"
)

func TestTableHeadsItsColumnsAndCarriesItsRows(t *testing.T) {
	assert.String(
		t,
		`<table class="numeric-1"><thead><tr><th>Alfa</th><th>Bravo</th></tr></thead><tbody><tr><td>charlie</td><td>1</td></tr></tbody></table>`,
		web_tester.Render(
			t,
			table.New("Alfa", "Bravo").
				Numeric(1).
				Rows(
					[]gomponents.Node{
						html.Tr(
							html.Td(gomponents.Text("charlie")),
							html.Td(gomponents.Text("1")),
						),
					},
				).
				Empty("Nothing yet.").
				Node()))
}

func TestNumericColumnsMarkTableBesideItsOwnClass(t *testing.T) {
	assert.StringContains(
		t,
		`<table class="echo numeric-0 numeric-3">`,
		web_tester.Render(
			t,
			table.New("Alfa").Class(constant.FixtureEcho).Numeric(3, 0).
				Node()))
}

func TestNumericColumnBeyondSharedRulesPanics(t *testing.T) {
	assert.PanicContains(
		t,
		"numeric column outside the shared rules",
		func() { table.New("Alfa").Numeric(constant.NumericColumnLimit) },
	)
}

func TestEveryNumericColumnHasSharedRule(t *testing.T) {
	for c := range constant.NumericColumnLimit {
		assert.StringContains(
			t,
			fmt.Sprintf(
				":where(table.numeric-%d) :is(th, td:not([colspan])):nth-child(%d)",
				c,
				c+1,
			),
			constant.ComponentStyle,
		)
	}
}

func TestEmptyTableSaysSoInSharedLook(t *testing.T) {
	assert.String(
		t,
		`<p class="absent">Nothing yet.</p>`,
		web_tester.Render(t, table.New("Alfa").Empty("Nothing yet.").Node()))
}

func TestTableWithoutEmptyMessageStaysTable(t *testing.T) {
	assert.String(
		t,
		`<table><thead><tr><th>Alfa</th></tr></thead><tbody></tbody></table>`,
		web_tester.Render(t, table.New("Alfa").Node()))
}

func TestHeaderlessTableCarriesOnlyItsBody(t *testing.T) {
	assert.String(
		t,
		`<table class="echo"><tbody><tr><td>alfa</td></tr></tbody></table>`,
		web_tester.Render(
			t,
			table.New().
				Headerless().
				Class(constant.FixtureEcho).
				Rows(
					[]gomponents.Node{html.Tr(html.Td(gomponents.Text("alfa")))},
				).
				Node()))
}

func TestIdentifiedTableKeepsItsTargetWhenEmpty(t *testing.T) {
	assert.String(
		t,
		`<div id="delta"><p class="absent">Nothing yet.</p></div>`,
		web_tester.Render(
			t,
			table.New("Alfa").Empty("Nothing yet.").Identifier("delta").
				Node()))
}

func TestTableCarriesItsClassAndBodyIdentifier(t *testing.T) {
	assert.String(
		t,
		`<table class="echo"><thead><tr><th>Alfa</th></tr></thead><tbody id="foxtrot"></tbody></table>`,
		web_tester.Render(
			t,
			table.New("Alfa").Class(constant.FixtureEcho).BodyIdentifier(
				"foxtrot",
			).
				Node()))
}
