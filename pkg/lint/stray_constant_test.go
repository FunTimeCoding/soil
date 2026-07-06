package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"strings"
	"testing"
)

func TestStrayConstantFlagged(t *testing.T) {
	l := StrayConstant(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nconst Foo = 1\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      constant.StrayConstantKey,
				Text:     constant.StrayConstantText,
				Path:     "Alfa",
				Type:     concern.Line,
				Line:     3,
				LineText: "const Foo = 1",
			},
		},
		"",
		l,
	)
}

func TestStrayConstantBlockFlagged(t *testing.T) {
	l := StrayConstant(
		upper.Bravo,
		strings.NewReader(
			"package example\n\nconst (\n\tFoo = 1\n\tBar = 2\n)\n",
		),
	)
	assertReport(
		t,
		"Bravo",
		true,
		[]*concern.Concern{
			{
				Key:      constant.StrayConstantKey,
				Text:     constant.StrayConstantText,
				Path:     "Bravo",
				Type:     concern.Line,
				Line:     3,
				LineText: "const (",
			},
		},
		"",
		l,
	)
}

func TestStrayConstantExemptByFilename(t *testing.T) {
	l := StrayConstant(
		"constant.go",
		strings.NewReader(
			"package example\n\nconst Foo = 1\n",
		),
	)
	assertReport(t, "constant.go", false, nil, "", l)
}

func TestStrayConstantExemptByFilenameTest(t *testing.T) {
	l := StrayConstant(
		"constant_test.go",
		strings.NewReader(
			"package example\n\nconst Foo = 1\n",
		),
	)
	assertReport(t, "constant_test.go", false, nil, "", l)
}

func TestStrayConstantExemptByPackage(t *testing.T) {
	l := StrayConstant(
		upper.Charlie,
		strings.NewReader(
			"package constant\n\nconst Foo = 1\n",
		),
	)
	assertReport(t, "Charlie", false, nil, "", l)
}

func TestStrayConstantExemptByConstantDirectory(t *testing.T) {
	l := StrayConstant(
		"pkg/tool/example/constant/color/item.go",
		strings.NewReader(
			"package color\n\nconst Foo = 1\n",
		),
	)
	assertReport(
		t,
		"pkg/tool/example/constant/color/item.go",
		false,
		nil,
		"",
		l,
	)
}

func TestStrayConstantDirectoryMatchesExactly(t *testing.T) {
	l := StrayConstant(
		"pkg/example/constants/foo.go",
		strings.NewReader(
			"package example\n\nconst Foo = 1\n",
		),
	)
	assertReport(
		t,
		"pkg/example/constants/foo.go",
		true,
		[]*concern.Concern{
			{
				Key:      constant.StrayConstantKey,
				Text:     constant.StrayConstantText,
				Path:     "pkg/example/constants/foo.go",
				Type:     concern.Line,
				Line:     3,
				LineText: "const Foo = 1",
			},
		},
		"",
		l,
	)
}

func TestStrayConstantInsideFunctionNotFlagged(t *testing.T) {
	l := StrayConstant(
		upper.Delta,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tconst x = 1\n\t_ = x\n}\n",
		),
	)
	assertReport(t, "Delta", false, nil, "", l)
}
