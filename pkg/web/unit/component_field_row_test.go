package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/field_row"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"testing"
)

func TestAFieldRowShowsItsLabelValueAndActions(t *testing.T) {
	assert.String(
		t,
		`<div id="field-alfa" class="field-row"><span class="field-label" data-tooltip="Bravo help">Bravo</span><span class="field-value">charlie</span><span class="field-actions"><button>delta</button></span></div>`,
		renderedNode(
			t,
			field_row.New("Bravo").
				Identifier("field-alfa").
				Tooltip("Bravo help").
				Value("charlie").
				Actions(html.Button(gomponents.Text("delta"))).
				Node(),
		),
	)
}

func TestAnEmptyFieldRowReadsNotSet(t *testing.T) {
	assert.String(
		t,
		`<div class="field-row field-unset"><span class="field-label">Echo</span><span class="field-value"><span class="absent">not set</span></span></div>`,
		renderedNode(t, field_row.New("Echo").Value("").Node()),
	)
}

func TestANoteIsMutedWithoutMarkingTheRowUnset(t *testing.T) {
	assert.String(
		t,
		`<div class="field-row"><span class="field-label">Foxtrot</span><span class="field-value"><small class="field-note">text · required</small></span></div>`,
		renderedNode(
			t,
			field_row.New("Foxtrot").Note("text · required").Node(),
		),
	)
}

func TestAnUnlabelledRowCarriesItsContent(t *testing.T) {
	assert.String(
		t,
		`<div class="field-row"><span class="field-value"><a href="/golf">Golf</a></span></div>`,
		renderedNode(
			t,
			field_row.New("").
				ValueNode(html.A(html.Href("/golf"), gomponents.Text("Golf"))).
				Node(),
		),
	)
}
