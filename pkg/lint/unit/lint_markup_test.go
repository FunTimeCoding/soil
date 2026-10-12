package unit

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
	"testing"
)

func TestMarkupClean(t *testing.T) {
	l := lint.Markup(
		stringsConstant.UpperBravo,
		strings.NewReader("---\nmyKey: myValue\n"),
	)
	assertReport(t, "Bravo", false, nil, "", l)
}

func TestMarkup(t *testing.T) {
	l := lint.Markup(
		stringsConstant.UpperAlfa,
		strings.NewReader("myKey: myValue\n"),
	)
	assertReport(
		t,
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
