package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/action"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestAnActionIsAFormThatWorksWithoutScript(t *testing.T) {
	assert.String(
		t,
		`<form class="action" method="post" action="/alfa"><input type="hidden" name="bravo" value="1"><button type="submit">Go</button></form>`,
		renderedAction(t, action.New("/alfa", "Go").Hidden("bravo", "1")),
	)
}

func TestASwappedActionIsUpgradedByHtmx(t *testing.T) {
	page := renderedAction(
		t,
		action.New("/alfa", "Go").Swap("#charlie", constant.SwapOuter).Confirm(
			"Sure?",
		),
	)
	assert.StringContains(t, `method="post" action="/alfa"`, page)
	assert.StringContains(t, `hx-post="/alfa"`, page)
	assert.StringContains(t, `hx-target="#charlie"`, page)
	assert.StringContains(t, `hx-swap="outerHTML"`, page)
	assert.StringContains(t, `hx-confirm="Sure?"`, page)
	assert.StringNotContains(t, "onsubmit", page)
}

func TestAPlainActionConfirmsBeforeSubmitting(t *testing.T) {
	page := renderedAction(t, action.New("/alfa", "Go").Confirm("Sure?"))
	assert.StringContains(t, `onsubmit="return confirm(&#34;Sure?&#34;)"`, page)
	assert.StringNotContains(t, "hx-confirm", page)
}

func TestASmallTitledActionNeedsNoLabel(t *testing.T) {
	page := renderedAction(t, action.New("/alfa", "").Small().Title("2 × 3"))
	assert.StringContains(t, `<form class="action small"`, page)
	assert.StringContains(
		t,
		`<button type="submit" title="2 × 3"></button>`,
		page,
	)
}

func TestAnActionCarriesItsClassesAndCanBeDisabled(t *testing.T) {
	page := renderedAction(
		t,
		action.New("/alfa", "Go").
			Class("delta").
			ButtonClass("outline").
			Disabled(true),
	)
	assert.StringContains(t, `<form class="action delta"`, page)
	assert.StringContains(
		t,
		`<button type="submit" class="outline" disabled>Go</button>`,
		page,
	)
}
