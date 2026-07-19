package unit_test

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"strings"
	"testing"
)

func TestFunctionEmptyBody(t *testing.T) {
	l := lint.Function(
		upper.Bravo,
		strings.NewReader(
			"package main\n\nfunc main() {\n}\n",
		),
	)
	assertReport(
		t,
		"Bravo",
		true,
		[]*concern.Concern{
			{
				Key:      "empty-function-body",
				Text:     "Function body with only whitespace",
				Path:     "Bravo",
				Type:     concern.Line,
				Line:     3,
				LineText: "func main() {\n}",
				Fixed:    true,
			},
		},
		"package main\n\nfunc main() {}\n",
		l,
	)
}

func TestFunctionClean(t *testing.T) {
	l := lint.Function(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tfmt.Println(\"hello\")\n}\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}
