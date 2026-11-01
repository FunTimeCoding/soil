package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/fact"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"testing"
)

func TestAFactTableLabelsEachRow(t *testing.T) {
	assert.String(
		t,
		`<table class="facts"><tbody><tr><th scope="row">Alfa</th><td>bravo</td></tr></tbody></table>`,
		renderedFact(t, fact.New().Row("Alfa", "bravo")),
	)
}

func TestASectionSpansBothColumns(t *testing.T) {
	assert.String(
		t,
		`<table class="facts"><tbody><tr class="section"><th colspan="2">Charlie</th></tr><tr><th scope="row">Alfa</th><td>bravo</td></tr></tbody></table>`,
		renderedFact(t, fact.New().Section("Charlie").Row("Alfa", "bravo")),
	)
}

func TestAFactValueCanBeMarkup(t *testing.T) {
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
