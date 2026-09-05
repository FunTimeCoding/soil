package unit

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
	"testing"
)

func TestMarkupClean(t *testing.T) {
	l := lint.Markup(
		constant.UpperBravo,
		strings.NewReader("---\nmyKey: myValue\n"),
	)
	assertReport(t, "Bravo", false, nil, "", l)
}

func TestMarkup(t *testing.T) {
	l := lint.Markup(constant.UpperAlfa, strings.NewReader("myKey: myValue\n"))
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      "front_matter_delimiter",
				Text:     "No front matter delimiter",
				Path:     "Alfa",
				Type:     lintConstant.ConcernLine,
				Line:     1,
				LineText: "myKey: myValue",
				Fixed:    true,
			},
		},
		"---\nmyKey: myValue\n",
		l,
	)
}
