package assert

import (
	"github.com/pmezard/go-difflib/difflib"
	"github.com/sanity-io/litter"
	"testing"
)

// Deep comparison over the exported surface only - unexported
// fields are invisible to the comparator, so black-box tests can
// whole-compare entities without reaching into private state.
func Exported(
	t *testing.T,
	expect any,
	actual any,
) {
	t.Helper()
	e := litter.Sdump(expect)
	a := litter.Sdump(actual)

	if e == a {
		return
	}

	text, f := difflib.GetUnifiedDiffString(
		difflib.UnifiedDiff{
			A:       difflib.SplitLines(e),
			B:       difflib.SplitLines(a),
			Context: 10,
		},
	)

	if f != nil {
		panic(f)
	}

	t.Error(text)
}
