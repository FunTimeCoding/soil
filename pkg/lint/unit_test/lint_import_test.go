package unit_test

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
	"testing"
)

func TestImportClean(t *testing.T) {
	l := lint.Import(
		constant.UpperBravo,
		strings.NewReader(
			"package example\n\nimport \"fmt\"\n\nfunc Example() {\n\tfmt.Println(\"Hello.\")\n}\n",
		),
	)
	assertReport(t, "Bravo", false, nil, "", l)
}

func TestImportBlank(t *testing.T) {
	l := lint.Import(
		constant.UpperCharlie,
		strings.NewReader(
			"package example\n\nimport (\n\t\"fmt\"\n\n\t\"log\"\n)\n\nfunc Example() {\n\tfmt.Printf(\"test\")\n\tlog.Printf(\"test\")\n}\n",
		),
	)
	assertReport(
		t,
		"Charlie",
		true,
		[]*concern.Concern{
			{
				Key:      "import_blank",
				Text:     "Import block contains blank line",
				Path:     "Charlie",
				Type:     lintConstant.ConcernLine,
				Line:     3,
				LineText: "import (",
				Fixed:    true,
			},
		},
		"package example\n\nimport (\n\t\"fmt\"\n\t\"log\"\n)\n\nfunc Example() {\n\tfmt.Printf(\"test\")\n\tlog.Printf(\"test\")\n}\n",
		l,
	)
}

func TestGo(t *testing.T) {
	l := lint.Import(
		constant.UpperAlfa,
		strings.NewReader(
			"package example\n\nimport (\n\t\"example.org/example/fmt\"\n)\n\nfunc Example() {\n\tfmt.Println(\"Hello friend.\")\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      "single_multi_import",
				Text:     "Single multi import",
				Path:     "Alfa",
				Type:     lintConstant.ConcernLine,
				Line:     3,
				LineText: "import (",
				Fixed:    true,
			},
		},
		"package example\n\nimport \"example.org/example/fmt\"\n\nfunc Example() {\n\tfmt.Println(\"Hello friend.\")\n}\n",
		l,
	)
}
