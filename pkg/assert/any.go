package assert

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/sanity-io/litter"
	"reflect"
	"testing"
)

func Any(
	t *testing.T,
	expect any,
	actual any,
) {
	t.Helper()

	if (expect == nil && actual != nil) ||
		(expect != nil && actual == nil) {
		t.Fatal(
			spew.Sdump(
				"\nNil comparison\nExpect: %#v\nActual: %#v",
				expect,
				actual,
			),
		)
	} else if expect == nil {
		return
	}

	expectLength := -1
	actualLength := -1

	switch reflect.TypeOf(expect).Kind() {
	case reflect.Slice:
		expectLength = reflect.ValueOf(expect).Len()
	case reflect.Array:
		expectLength = reflect.ValueOf(expect).Len()
	default:
		// pass
	}

	switch reflect.TypeOf(actual).Kind() {
	case reflect.Slice:
		actualLength = reflect.ValueOf(actual).Len()
	case reflect.Array:
		actualLength = reflect.ValueOf(actual).Len()
	default:
		// pass
	}

	if expectLength > -1 || actualLength > -1 {
		var reason string

		if actualLength < expectLength {
			reason = "Actual is shorter than expect"
		} else if actualLength > expectLength {
			reason = "Actual is longer than expect"
		}

		if reason != "" {
			diff := difflib.UnifiedDiff{
				A:        difflib.SplitLines(litter.Sdump(expect)),
				B:        difflib.SplitLines(litter.Sdump(actual)),
				FromFile: fmt.Sprintf("Expect (%d)", expectLength),
				ToFile:   fmt.Sprintf("Actual (%d)", actualLength),
				Context:  1,
			}
			text, e := difflib.GetUnifiedDiffString(diff)

			if e != nil {
				panic(e)
			}

			t.Errorf("%s\n%s", reason, text)
		}
	}

	if reflect.DeepEqual(actual, expect) {
		return
	}

	text, f := difflib.GetUnifiedDiffString(
		difflib.UnifiedDiff{
			A:       difflib.SplitLines(spew.Sdump(expect)),
			B:       difflib.SplitLines(spew.Sdump(actual)),
			Context: 10,
		},
	)

	if f != nil {
		panic(f)
	}

	t.Error(text)
}
