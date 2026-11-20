package unit

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/unit/lint_tester"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
	"testing"
)

func TestMarkupClean(t *testing.T) {
	l := lint.Markup(
		stringsConstant.UpperBravo,
		strings.NewReader("---\nmyKey: myValue\n"),
	)
	lint_tester.New(t).Report("Bravo", false, nil, "", l)
}

func TestMarkup(t *testing.T) {
	l := lint.Markup(
		stringsConstant.UpperAlfa,
		strings.NewReader("myKey: myValue\n"),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      "front_matter_delimiter",
				Text:     "No front matter delimiter",
				Path:     "Alfa",
				Type:     constant.ConcernLine,
				Line:     1,
				LineText: "myKey: myValue",
				Fixed:    true,
			},
		},
		"---\nmyKey: myValue\n",
		l,
	)
}
