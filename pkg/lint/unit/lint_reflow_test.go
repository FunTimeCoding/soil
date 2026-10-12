package unit

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
	"testing"
)

func TestReflowClean(t *testing.T) {
	l := lint.Reflow(stringsConstant.UpperAlfa, strings.NewReader("one two\n"))
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestReflowRewrapsAnOverrunningParagraph(t *testing.T) {
	content := join.Empty(strings.Repeat("word ", 17), "end\n")
	l := lint.Reflow(stringsConstant.UpperAlfa, strings.NewReader(content))
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      "reflow",
				Text:     "Line exceeds the column width",
				Path:     "Alfa",
				Type:     constant.ConcernLine,
				Line:     1,
				LineText: strings.TrimSuffix(content, "\n"),
				Fixed:    true,
			},
		},
		"word word word word word word word word word word word word word word word word\nword end\n",
		l,
	)
}

func TestReflowRefusesAHardLineBreak(t *testing.T) {
	content := join.Empty(strings.Repeat("word ", 17), "end  \nnext\n")
	l := lint.Reflow(stringsConstant.UpperAlfa, strings.NewReader(content))
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:   "reflow_refused",
				Text:  "Reflow refused - document structure changed",
				Path:  "Alfa",
				Type:  constant.ConcernFile,
				Fixed: false,
			},
		},
		"",
		l,
	)
}
