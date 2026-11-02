package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/pointer_tester"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
	"testing"
)

func TestPointersSymbolDeclared(t *testing.T) {
	l := pointer_tester.Declaring(
		[]string{"pkg/provision/salt.Client", "pkg/system/run.New"},
		"pkg/provision/salt",
		"pkg/system/run",
	)(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Wraps `go:pkg/provision/salt.Client` and `go:pkg/system/run/New()`.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersSymbolUndeclared(t *testing.T) {
	line := "Wraps `go:pkg/provision/salt.Clint` today."
	l := pointer_tester.Declaring(
		[]string{"pkg/provision/salt.Client"},
		"pkg/provision/salt",
	)(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(
		t,
		"Alfa",
		true,
		pointer_tester.Undeclared("Alfa", line),
		"",
		l,
	)
}

func TestPointersSymbolMemberChecksItsHead(t *testing.T) {
	l := pointer_tester.Declaring(
		[]string{"pkg/provision/salt.Client"},
		"pkg/provision/salt",
	)(
		stringsConstant.UpperAlfa,
		strings.NewReader("Runs `go:pkg/provision/salt.Client.Run()` once.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}
