package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/report"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/revision"
	"testing"
)

func TestReportSeparatesBodyAndSignatureChanges(t *testing.T) {
	r := report.New(
		revision.New(
			map[string][]byte{
				"go.mod": []byte("module example.org/root\n"),
				"alfa.go": []byte(
					"package alfa\n\nfunc Bravo() int {\n\treturn 1\n}\n\nfunc Charlie(a int) int {\n\treturn a\n}\n\nfunc Delta() {}\n",
				),
			},
		),
		revision.New(
			map[string][]byte{
				"go.mod": []byte("module example.org/root\n"),
				"alfa.go": []byte(
					"package alfa\n\nfunc Bravo() int {\n\treturn 2\n}\n\nfunc Charlie(a int, b int) int {\n\treturn a\n}\n\nfunc Delta() {}\n",
				),
			},
		),
	)
	assert.Strings(
		t,
		[]string{
			"unchanged . alfa Delta",
			"signature . alfa Charlie",
			"body . alfa Bravo",
		},
		reportLines(r),
	)
}

func TestReportFollowsAPackageMovedWhole(t *testing.T) {
	r := report.New(
		revision.New(
			map[string][]byte{
				"go.mod": []byte("module example.org/root\n"),
				"alfa/alfa.go": []byte(
					"package alfa\n\ntype Card struct{}\n\nfunc New() *Card {\n\treturn &Card{}\n}\n",
				),
				"charlie/charlie.go": []byte(
					"package charlie\n\nimport \"example.org/root/alfa\"\n\nfunc Make() *alfa.Card {\n\treturn alfa.New()\n}\n",
				),
			},
		),
		revision.New(
			map[string][]byte{
				"go.mod": []byte("module example.org/root\n"),
				"bravo/bravo.go": []byte(
					"package bravo\n\ntype Card struct{}\n\nfunc New() *Card {\n\treturn &Card{}\n}\n",
				),
				"charlie/charlie.go": []byte(
					"package charlie\n\nimport \"example.org/root/bravo\"\n\nfunc Make() *bravo.Card {\n\treturn bravo.New()\n}\n",
				),
			},
		),
	)
	assert.Strings(
		t,
		[]string{
			"unchanged charlie charlie Make",
			"moved bravo bravo Card",
			"moved bravo bravo New",
		},
		reportLines(r),
	)
}

func TestReportFindsARenameByShape(t *testing.T) {
	r := report.New(
		revision.New(
			map[string][]byte{
				"go.mod": []byte("module example.org/root\n"),
				"alfa.go": []byte(
					"package alfa\n\nfunc Sum(a int, b int) int {\n\treturn a + b\n}\n",
				),
			},
		),
		revision.New(
			map[string][]byte{
				"go.mod": []byte("module example.org/root\n"),
				"alfa.go": []byte(
					"package alfa\n\nfunc Total(a int, b int) int {\n\treturn a + b\n}\n",
				),
			},
		),
	)
	assert.Strings(t, []string{"renamed . alfa Total"}, reportLines(r))
}

func TestReportMarksAMoveThatChanged(t *testing.T) {
	r := report.New(
		revision.New(
			map[string][]byte{
				"go.mod": []byte("module example.org/root\n"),
				"alfa/limit.go": []byte(
					"package alfa\n\nfunc Limit() int {\n\treturn 3\n}\n",
				),
			},
		),
		revision.New(
			map[string][]byte{
				"go.mod": []byte("module example.org/root\n"),
				"bravo/limit.go": []byte(
					"package bravo\n\nfunc Limit() int {\n\treturn 4\n}\n",
				),
			},
		),
	)
	assert.Strings(
		t,
		[]string{"moved-changed bravo bravo Limit"},
		reportLines(r),
	)
}

func TestReportListsAddedAndRemoved(t *testing.T) {
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
		[]string{"added . alfa Foxtrot", "removed . alfa Echo"},
		reportLines(r),
	)
}
