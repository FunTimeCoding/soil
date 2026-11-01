package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/choice"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"github.com/funtimecoding/soil/pkg/web/component/entry/field"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"testing"
)

func TestAnEntryIsALabelledFormThatWorksWithoutScript(t *testing.T) {
	assert.String(
		t,
		`<form class="entry" method="post" action="/alfa"><input type="hidden" name="bravo" value="1"><label><span class="entry-label">Charlie</span><input type="text" name="charlie" value="delta" placeholder="echo" required></label><button type="submit">Go</button></form>`,
		renderedNode(
			t,
			entry.New("/alfa", "Go").
				Hidden("bravo", "1").
				Text("charlie", "Charlie", "delta").
				Placeholder(constant.FixtureEcho).
				Required().
				Node(),
		),
	)
}

func TestAFileFieldMakesTheEntryAMultipartForm(t *testing.T) {
	assert.String(
		t,
		`<form class="entry" method="post" action="/alfa" enctype="multipart/form-data"><label><span class="entry-label">Bravo</span><input type="file" name="bravo" accept="image/*"></label><button type="submit">Go</button></form>`,
		renderedNode(
			t,
			entry.New("/alfa", "Go").File("bravo", "Bravo", "image/*").Node(),
		),
	)
}

func TestASwappedInlineEntryCarriesItsAlertAndIdentifier(t *testing.T) {
	page := renderedNode(
		t,
		entry.New("/alfa", "Go").
			Swap("#foxtrot", constant.SwapOuter).
			Identifier("foxtrot").
			Inline().
			Alert("Not like that.").
			Node(),
	)
	assert.StringContains(
		t,
		`<form class="entry inline" id="foxtrot" method="post" action="/alfa" hx-post="/alfa" hx-target="#foxtrot" hx-swap="outerHTML"><p role="alert" class="form-alert">Not like that.</p>`,
		page,
	)
}

func TestAQueryEntryReadsWithGetAndSwapsByGet(t *testing.T) {
	assert.StringContains(
		t,
		`<form class="entry inline" method="get" action="/alfa" hx-get="/alfa" hx-target="#foxtrot" hx-swap="innerHTML">`,
		renderedNode(
			t,
			entry.New("/alfa", "Filter").
				Query().
				Swap("#foxtrot", constant.SwapInner).
				Inline().
				Node(),
		),
	)
}

func TestAPasswordNeverCarriesAValueBesideAMail(t *testing.T) {
	assert.StringContains(
		t,
		`<input type="email" name="bravo" value="charlie@delta.test"></label><label><span class="entry-label">Golf</span><input type="password" name="golf"></label>`,
		renderedNode(
			t,
			entry.New("/alfa", "Go").
				Mail("bravo", "Bravo", "charlie@delta.test").
				Password("golf", "Golf").
				Node(),
		),
	)
}

func TestASearchFieldCanTakeTheFocus(t *testing.T) {
	assert.StringContains(
		t,
		`<input type="search" name="bravo" value="charlie" autofocus>`,
		renderedNode(
			t,
			entry.New("/alfa", "Go").
				Query().
				Search("bravo", "Bravo", "charlie").
				Focus().
				Node(),
		),
	)
}

func TestAHeadingSitsInsideTheFormAboveItsAlert(t *testing.T) {
	assert.StringContains(
		t,
		`action="/alfa"><h5>Sierra</h5><p role="alert" class="form-alert">Tango</p>`,
		renderedNode(
			t,
			entry.New("/alfa", "Go").Heading("Sierra").Alert("Tango").Node(),
		),
	)
}

func TestGroupsGatherTheFieldsThatFollowThem(t *testing.T) {
	page := renderedNode(
		t,
		entry.New("/alfa", "Go").
			Text("uniform", "Uniform", "").
			Group("Victor").
			Range("whiskey", "Whiskey", "0.5").
			Min("0").
			Max("1").
			Step("0.05").
			Content(html.P(gomponents.Text("xray"))).
			Group("Yankee").
			Text("zulu", "Zulu", "").
			Node(),
	)
	assert.StringContains(
		t,
		`<label><span class="entry-label">Uniform</span><input type="text" name="uniform"></label><fieldset><legend>Victor</legend><label><span class="entry-label">Whiskey</span><input type="range" name="whiskey" value="0.5" min="0" max="1" step="0.05"></label><p>xray</p></fieldset><fieldset><legend>Yankee</legend>`,
		page,
	)
}

func TestEndingAGroupReturnsToTheFormItself(t *testing.T) {
	assert.StringContains(
		t,
		`<fieldset><legend>Victor</legend><label><span class="entry-label">Uniform</span><input type="text" name="uniform"></label></fieldset><p>xray</p><button`,
		renderedNode(
			t,
			entry.New("/alfa", "Go").
				Group("Victor").
				Text("uniform", "Uniform", "").
				EndGroup().
				Content(html.P(gomponents.Text("xray"))).
				Node(),
		),
	)
}

func TestAPreviewPostsChangesElsewhereWhileTheSubmitStaysReal(t *testing.T) {
	assert.StringContains(
		t,
		`method="post" action="/alfa" hx-post="/alfa/preview" hx-trigger="change delay:300ms" hx-target="#card" hx-swap="outerHTML"`,
		renderedNode(
			t,
			entry.New("/alfa", "Go").Preview("/alfa/preview", "#card").Node(),
		),
	)
}

func TestTheButtonAndAFieldCarryTheirOwnClasses(t *testing.T) {
	page := renderedNode(
		t,
		entry.New("/alfa", "Go").
			ButtonClass("outline").
			Text("bravo", "Bravo", "").
			FieldClass("wide").
			Node(),
	)
	assert.StringContains(
		t,
		`<input type="text" name="bravo" class="wide">`,
		page,
	)
	assert.StringContains(
		t,
		`<button type="submit" class="outline">Go</button>`,
		page,
	)
}

func TestARecordEntryKeepsItsButtonsSideBySide(t *testing.T) {
	assert.String(
		t,
		`<form class="entry record" method="post" action="/alfa"><label><span class="entry-label">Bravo</span><input type="text" name="bravo"></label><div class="entry-buttons"><button type="submit">Go</button><a href="/charlie">Cancel</a></div></form>`,
		renderedNode(
			t,
			entry.New("/alfa", "Go").
				Record().
				Text("bravo", "Bravo", "").
				Beside(html.A(html.Href("/charlie"), gomponents.Text("Cancel"))).
				Node(),
		),
	)
}

func TestARevealingFieldFetchesTheNextFieldsWhenItChanges(t *testing.T) {
	assert.StringContains(
		t,
		`<select name="delta" required hx-get="/echo" hx-trigger="input changed delay:300ms" hx-target="#foxtrot" hx-swap="innerHTML" hx-include="closest form">`,
		renderedNode(
			t,
			entry.New("/alfa", "Go").
				Select("delta", "Delta", nil, "").
				Required().
				Reveal(constant.FixtureEchoPath, "#foxtrot").
				Node(),
		),
	)
}

func TestAnAlternativePostsTheSameFieldsElsewhereAfterThePrimary(t *testing.T) {
	assert.StringContains(
		t,
		`<div class="entry-buttons"><button type="submit">Go</button><button type="submit" class="entry-alternative" formaction="/skip" formnovalidate>Skip</button></div>`,
		renderedNode(
			t,
			entry.New("/alfa", "Go").
				Text("bravo", "Bravo", "").
				Required().
				Alternative("/skip", "Skip").
				Node(),
		),
	)
}

func TestSuggestionsOfferAListBesideTheField(t *testing.T) {
	assert.StringContains(
		t,
		`<label><span class="entry-label">Bravo</span><input type="text" name="bravo" list="bravo-suggestions"><datalist id="bravo-suggestions"><option value="charlie"></option><option value="delta"></option></datalist></label>`,
		renderedNode(
			t,
			entry.New("/alfa", "Go").
				Text("bravo", "Bravo", "").
				Suggestions([]string{"charlie", "delta"}).
				Node(),
		),
	)
}

func TestAFieldBuiltElsewhereJoinsTheEntryAndRendersAlone(t *testing.T) {
	f := field.New(constant.EntryNumberKind, "golf", "Golf", "3")
	assert.String(
		t,
		`<label><span class="entry-label">Golf</span><input type="number" name="golf" value="3"></label>`,
		renderedNode(t, f.Node()),
	)
	assert.StringContains(
		t,
		`action="/alfa"><label><span class="entry-label">Golf</span>`,
		renderedNode(t, entry.New("/alfa", "Go").Field(f).Node()),
	)
}

func TestAFieldChainedAloneMatchesTheEntryModifiers(t *testing.T) {
	chained := renderedNode(
		t,
		entry.New("/alfa", "Go").
			Field(
				field.New(constant.EntryNumberKind, "hotel", "Hotel", "2").
					Required().
					Min("1").
					Max("4").
					Step("1").
					Class("wide").
					Reveal("/india", "#juliett"),
			).
			Node(),
	)
	modified := renderedNode(
		t,
		entry.New("/alfa", "Go").
			Number("hotel", "Hotel", "2").
			Required().
			Min("1").
			Max("4").
			Step("1").
			FieldClass("wide").
			Reveal("/india", "#juliett").
			Node(),
	)
	assert.String(t, modified, chained)
	assert.StringContains(
		t,
		`<input type="number" name="hotel" value="2" required min="1" max="4" step="1" class="wide" hx-get="/india"`,
		chained,
	)
}

func TestARequiredSelectPromptsWithADisabledEmptyOption(t *testing.T) {
	page := renderedNode(
		t,
		entry.New("/alfa", "Go").
			Select(
				"golf",
				"Golf",
				[]*choice.Choice{
					choice.New("hotel", "Hotel"),
					choice.New("india", "India"),
				},
				"india",
			).
			Prompt("- pick one -").
			Required().
			Node(),
	)
	assert.StringContains(
		t,
		`<select name="golf" required><option value="" disabled>- pick one -</option><option value="hotel">Hotel</option><option value="india" selected>India</option></select>`,
		page,
	)
}

func TestAnOptionalSelectLeavesItsPromptSelectable(t *testing.T) {
	page := renderedNode(
		t,
		entry.New("/alfa", "Go").
			Select("golf", "Golf", nil, "").
			Prompt("- unknown -").
			Node(),
	)
	assert.StringContains(
		t,
		`<option value="" selected>- unknown -</option>`,
		page,
	)
}

func TestAnAppendingEntryResetsAfterItsRequest(t *testing.T) {
	page := renderedNode(
		t,
		entry.New("/alfa", "Add").
			Swap("#lima", constant.SwapAppend).
			ResetAfter().
			Link("mike", "Mike", "").
			Date("november", "November", "2026-10-26").
			TextArea("oscar", "Oscar", "papa").
			Lines("3").
			Select("quebec", "Quebec", choice.FromValues([]string{"romeo"}), "").
			Node(),
	)
	assert.StringContains(t, `hx-on::after-request="this.reset()"`, page)
	assert.StringContains(t, `<input type="url" name="mike">`, page)
	assert.StringContains(
		t,
		`<input type="date" name="november" value="2026-10-26">`,
		page,
	)
	assert.StringContains(
		t,
		`<textarea name="oscar" rows="3">papa</textarea>`,
		page,
	)
	assert.StringContains(t, `<option value="romeo">romeo</option>`, page)
}

func TestNumbersCarryTheirBoundsAndCheckboxesTheirState(t *testing.T) {
	page := renderedNode(
		t,
		entry.New("/alfa", "Go").
			Number("juliett", "Juliett", "2").
			Min("1").
			Max("4").
			Step("1").
			Checkbox("kilo", "Kilo", true).
			Node(),
	)
	assert.StringContains(
		t,
		`<input type="number" name="juliett" value="2" min="1" max="4" step="1">`,
		page,
	)
	assert.StringContains(
		t,
		`<label><input type="checkbox" name="kilo" value="on" checked>Kilo</label>`,
		page,
	)
}
