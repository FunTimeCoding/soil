package unit

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
	"testing"
)

func TestFixtureDirectiveSanctionsStrayConstant(t *testing.T) {
	path := "pkg/example/testdata/src/inner/machinery_test.go"
	l := lint.StrayConstant(
		path,
		strings.NewReader(constant.FixtureSanctionedSample),
	)
	assertReport(t, path, false, nil, "", l)
}

func TestFixtureDirectiveValidInTestdata(t *testing.T) {
	path := "pkg/example/testdata/src/inner/machinery_test.go"
	l := lint.FixtureDirective(
		path,
		strings.NewReader(constant.FixtureSanctionedSample),
	)
	assertReport(t, path, false, nil, "", l)
}

func TestFixtureDirectiveOutsideTestdata(t *testing.T) {
	path := "pkg/example/machinery.go"
	l := lint.FixtureDirective(
		path,
		strings.NewReader(constant.FixtureSanctionedSample),
	)
	assertReport(
		t,
		path,
		true,
		[]*concern.Concern{
			{
				Key:      constant.FixtureKey,
				Text:     constant.FixtureOutsideTestdataText,
				Path:     path,
				Type:     constant.ConcernLine,
				Line:     3,
				LineText: "// golint:fixture stray_constant",
			},
		},
		"",
		l,
	)
}

func TestFixtureDirectiveOutsideTestdataDoesNotSanction(t *testing.T) {
	path := "pkg/example/machinery.go"
	l := lint.StrayConstant(
		path,
		strings.NewReader(constant.FixtureSanctionedSample),
	)
	assertReport(
		t,
		path,
		true,
		[]*concern.Concern{
			{
				Key:      constant.StrayConstantKey,
				Text:     constant.StrayConstantText,
				Path:     path,
				Type:     constant.ConcernLine,
				Line:     4,
				LineText: "const Foo = 1",
			},
		},
		"",
		l,
	)
}

func TestFixtureDirectiveUnknownRule(t *testing.T) {
	path := "pkg/example/testdata/src/inner/machinery.go"
	l := lint.FixtureDirective(
		path,
		strings.NewReader(
			"// golint:fixture unknown_rule\n\npackage example\n",
		),
	)
	assertReport(
		t,
		path,
		true,
		[]*concern.Concern{
			{
				Key:      constant.FixtureKey,
				Text:     constant.FixtureUnknownRuleText,
				Path:     path,
				Type:     constant.ConcernLine,
				Line:     1,
				LineText: "// golint:fixture unknown_rule",
			},
		},
		"",
		l,
	)
}

func TestFixtureDirectiveMisplaced(t *testing.T) {
	path := "pkg/example/testdata/src/inner/machinery.go"
	l := lint.FixtureDirective(
		path,
		strings.NewReader(constant.FixtureMisplacedSample),
	)
	assertReport(
		t,
		path,
		true,
		[]*concern.Concern{
			{
				Key:      constant.FixtureKey,
				Text:     constant.FixtureMisplacedText,
				Path:     path,
				Type:     constant.ConcernLine,
				Line:     3,
				LineText: "// golint:fixture stray_constant",
			},
		},
		"",
		l,
	)
}

func TestFixtureDirectiveDangling(t *testing.T) {
	path := "pkg/example/testdata/src/inner/machinery.go"
	l := lint.FixtureDirective(
		path,
		strings.NewReader(constant.FixtureDanglingSample),
	)
	assertReport(
		t,
		path,
		true,
		[]*concern.Concern{
			{
				Key:      constant.FixtureKey,
				Text:     constant.FixtureMisplacedText,
				Path:     path,
				Type:     constant.ConcernLine,
				Line:     3,
				LineText: "// golint:fixture stray_constant",
			},
		},
		"",
		l,
	)
}

func TestFixtureDirectiveMisplacedDoesNotSanction(t *testing.T) {
	path := "pkg/example/testdata/src/inner/machinery.go"
	l := lint.StrayConstant(
		path,
		strings.NewReader(constant.FixtureMisplacedSample),
	)
	assertReport(
		t,
		path,
		true,
		[]*concern.Concern{
			{
				Key:      constant.StrayConstantKey,
				Text:     constant.StrayConstantText,
				Path:     path,
				Type:     constant.ConcernLine,
				Line:     5,
				LineText: "const Foo = 1",
			},
		},
		"",
		l,
	)
}
