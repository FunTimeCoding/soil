package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/pointer_tester"
	"github.com/funtimecoding/soil/pkg/lint/unit/lint_tester"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
	"testing"
)

func TestPointersMissingTestNameUnchecked(t *testing.T) {
	checker, seen := pointer_tester.Symbols(
		nil,
		nil,
		[]string{"TestShelfHoldsStock"},
	)
	l := checker(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Pinned by `TestShelfHoldsStock` and `TestShelfLostStock()`, shaped like `TestXxx`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
	assert.Strings(t, []string{"Alfa:1 test_name TestShelfLostStock()"}, *seen)
}

func TestPointersUndeclaredMemberUnchecked(t *testing.T) {
	checker, seen := pointer_tester.Symbols(
		[]string{
			"pkg/shop/store.Store",
			"pkg/shop/store.Store.Open",
			"pkg/shop/store.Default",
			"pkg/shop/store.Shelf",
			"pkg/shop/store.Shelf.<any>",
		},
		[]string{"pkg/shop/store.Store", "pkg/shop/store.Shelf"},
		nil,
		"pkg/shop/store",
	)
	l := checker(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Opens `go:pkg/shop/store.Store.Open()`, closes `go:pkg/shop/store.Store.Close()`, names `go:pkg/shop/store.Default.Name` and locks `go:pkg/shop/store.Shelf.Lock()`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
	assert.Strings(
		t,
		[]string{"Alfa:1 member go:pkg/shop/store.Store.Close()"},
		*seen,
	)
}

func TestPointersBareSymbolUnchecked(t *testing.T) {
	checker, seen := pointer_tester.Symbols(nil, nil, nil)
	l := checker(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Calls `store.Open`, `Store.Close()` and `store.Store.Open`, beside `main.go`, `README.md`, `Notes.md` and `h.example.org`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
	assert.Strings(
		t,
		[]string{
			"Alfa:1 bare_symbol store.Open",
			"Alfa:1 bare_symbol Store.Close()",
			"Alfa:1 bare_symbol store.Store.Open",
		},
		*seen,
	)
}
