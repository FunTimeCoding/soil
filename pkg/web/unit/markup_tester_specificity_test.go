package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/unit/markup_tester"
	"testing"
)

func TestSpecificityCountsWhereAsNothingAndIsAsItsWeightiestArgument(t *testing.T) {
	weigh := func(
		expected [3]int,
		selector string,
	) {
		t.Helper()
		assert.Any(t, expected, markup_tester.Specificity(selector))
	}
	weigh([3]int{0, 0, 1}, "td")
	weigh([3]int{0, 1, 0}, ":where(table) .numeric")
	weigh([3]int{0, 0, 0}, ":where(table .numeric)")
	weigh([3]int{0, 1, 0}, ":where(p).empty")
	weigh([3]int{0, 2, 0}, ".field-actions :is(button, [role=button])")
	weigh(
		[3]int{0, 3, 1},
		":where(form).entry.inline :is(input:not([type=checkbox]), select, textarea)",
	)
	weigh(
		[3]int{0, 3, 1},
		":where(form).entry.record label:has(> .entry-label)",
	)
	weigh([3]int{1, 1, 0}, "#notifications:empty")
	weigh([3]int{0, 0, 2}, "dialog::backdrop")
}

func TestAlternativesUnfoldWhereAndIsAndDropStates(t *testing.T) {
	assert.Any(
		t,
		[]string{"table .numeric"},
		markup_tester.Alternatives(":where(table) .numeric"),
	)
	assert.Any(
		t,
		[]string{".field-actions button", ".field-actions [role=button]"},
		markup_tester.Alternatives(".field-actions :is(button, [role=button])"),
	)
	assert.Any(t, []string{"a"}, markup_tester.Alternatives("a:hover"))
	assert.Any(
		t,
		[]string{"label:has(.entry-label)"},
		markup_tester.Alternatives("label:has(> .entry-label)"),
	)
}
