package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/style_sheet"
	"testing"
)

func TestARuleCarriesItsSelectorsDeclarationsAndOffset(t *testing.T) {
	rules := style_sheet.Parse(
		"\n.alfa, :where(.bravo, .charlie) > p { color:  red;\n margin: 0 }",
	)
	assert.Count(t, 1, rules)
	assert.Integer(t, 1, rules[0].Offset)
	assert.Strings(
		t,
		[]string{".alfa", ":where(.bravo, .charlie) > p"},
		rules[0].Selectors,
	)
	assert.Strings(
		t,
		[]string{"color: red", "margin: 0"},
		rules[0].Declarations,
	)
}

func TestCommentsAndKeyframesAreNotRulesButMediaRulesAre(t *testing.T) {
	rules := style_sheet.Parse(
		"/* .delta { x: y } */ @keyframes echo { from { a: b } to { a: c } } @media (min-width: 1px) { .foxtrot { a: b } } .golf { a: b }",
	)
	assert.Count(t, 2, rules)
	assert.Strings(t, []string{".foxtrot"}, rules[0].Selectors)
	assert.Strings(t, []string{".golf"}, rules[1].Selectors)
}

func TestSplitTopKeepsBracketsWhole(t *testing.T) {
	assert.Strings(
		t,
		[]string{"a[b=\"c,d\"]", ":is(e, f)", "g"},
		style_sheet.SplitTop("a[b=\"c,d\"], :is(e, f), g", ','),
	)
}
