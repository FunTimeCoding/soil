package unit

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
	"testing"
)

func TestVariableErrorAssignment(t *testing.T) {
	l := lint.Variable(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\terr := foo()\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      constant.ErrorVariableKey,
				Text:     "Use e instead of err for error variable",
				Path:     "Alfa",
				Type:     constant.ConcernLine,
				Line:     4,
				LineText: "\terr := foo()",
			},
		},
		"",
		l,
	)
}

func TestVariableErrorMultiReturn(t *testing.T) {
	l := lint.Variable(
		stringsConstant.UpperBravo,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx, err := foo()\n}\n",
		),
	)
	assertReport(
		t,
		"Bravo",
		true,
		[]*concern.Concern{
			{
				Key:      constant.ErrorVariableKey,
				Text:     "Use e instead of err for error variable",
				Path:     "Bravo",
				Type:     constant.ConcernLine,
				Line:     4,
				LineText: "\tx, err := foo()",
			},
		},
		"",
		l,
	)
}

func TestVariableErrorComparison(t *testing.T) {
	l := lint.Variable(
		stringsConstant.UpperCharlie,
		strings.NewReader(
			"package example\n\nfunc Example() bool {\n\treturn err == nil\n}\n",
		),
	)
	assertReport(t, "Charlie", false, nil, "", l)
}

func TestVariableEOkay(t *testing.T) {
	l := lint.Variable(
		stringsConstant.UpperDelta,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\te := foo()\n}\n",
		),
	)
	assertReport(t, "Delta", false, nil, "", l)
}

func TestVariableErrorInString(t *testing.T) {
	l := lint.Variable(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\ts := \"err := foo()\"\n\t_ = s\n}\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestVariableErrorComment(t *testing.T) {
	l := lint.Variable(
		stringsConstant.UpperAlfa,
		strings.NewReader("package example\n\n// err := foo()\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}
