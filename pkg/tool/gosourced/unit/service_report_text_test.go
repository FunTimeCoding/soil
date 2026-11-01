package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/report"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/revision"
	"testing"
)

func TestReportTextKeepsTheSourceAsWritten(t *testing.T) {
	r := revision.New(
		map[string][]byte{
			"go.mod": []byte("module example.org/root\n"),
			"alfa.go": []byte(
				"package alfa\n\nconst Bravo = 1\n\nconst Charlie = Bravo\n\nfunc Delta() int {\n\t// echo\n\treturn Charlie\n}\n",
			),
		},
	)
	assert.String(
		t,
		"Charlie = Bravo",
		revisionSymbol(t, r, ". alfa Charlie").Text,
	)
	assert.String(
		t,
		"func Delta() int {\n\t// echo\n\treturn Charlie\n}",
		revisionSymbol(t, r, ". alfa Delta").Text,
	)
}

func TestReportTextShowsAddedWholeAndRemovedByItsFirstLine(t *testing.T) {
	r := report.New(
		revision.New(
			map[string][]byte{
				"go.mod": []byte("module example.org/root\n"),
				"alfa.go": []byte(
					"package alfa\n\nfunc Echo() int {\n\treturn 1\n}\n",
				),
			},
		),
		revision.New(
			map[string][]byte{
				"go.mod": []byte("module example.org/root\n"),
				"alfa.go": []byte(
					"package alfa\n\nfunc Foxtrot() string {\n\treturn \"golf\"\n}\n",
				),
			},
		),
	)
	assert.Strings(
		t,
		[]string{"func Foxtrot() string {", "\treturn \"golf\"", "}"},
		report.EntryText(r.Entries[0]),
	)
	assert.Strings(
		t,
		[]string{"func Echo() int {"},
		report.EntryText(r.Entries[1]),
	)
}
