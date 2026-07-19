package unit_test

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"strings"
	"testing"
)

func TestMarkupClean(t *testing.T) {
	l := lint.Markup(
		upper.Bravo,
		strings.NewReader("---\nmyKey: myValue\n"),
	)
	assertReport(t, "Bravo", false, nil, "", l)
}

func TestMarkup(t *testing.T) {
	l := lint.Markup(
		upper.Alfa,
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
				Type:     concern.Line,
				Line:     1,
				LineText: "myKey: myValue",
				Fixed:    true,
			},
		},
		"---\nmyKey: myValue\n",
		l,
	)
}
