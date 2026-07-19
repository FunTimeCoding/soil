package unit_test

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"strings"
	"testing"
)

func pointerExists(existing []string) func(string) bool {
	return func(p string) bool {
		for _, e := range existing {
			if p == e {
				return true
			}
		}

		return false
	}
}

func pointerChecker(existing ...string) lint.Checker {
	return lint.Pointers(
		[]string{".claude", "doc", "pkg"},
		pointerExists(existing),
		pointerExists(existing),
		func(string) bool { return false },
	)
}

func gitignoredChecker(ignored ...string) lint.Checker {
	return lint.Pointers(
		[]string{".claude", "doc", "pkg"},
		func(string) bool { return false },
		func(string) bool { return false },
		pointerExists(ignored),
	)
}

func TestPointersGitignored(t *testing.T) {
	l := gitignoredChecker(".claude/notes/alfa.md")(
		upper.Alfa,
		strings.NewReader(
			"Style guide at `.claude/notes/alfa.md` today.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersClean(t *testing.T) {
	l := pointerChecker()(
		upper.Alfa,
		strings.NewReader("Run `task lint` now.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersExisting(t *testing.T) {
	l := pointerChecker("doc/ai/runbook/lint.md")(
		upper.Alfa,
		strings.NewReader("Read `doc/ai/runbook/lint.md` first.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersDirectory(t *testing.T) {
	l := pointerChecker("doc/ai/spec")(
		upper.Alfa,
		strings.NewReader("Specs live in `doc/ai/spec/`.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersIgnored(t *testing.T) {
	l := pointerChecker()(
		upper.Alfa,
		strings.NewReader(
			"See `tmp/gosec.json`, `doc/ai/runbook/<name>.md`, and https://example.org/page.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersDead(t *testing.T) {
	line := "Read `doc/ai/runbook/ghost.md` first."
	l := pointerChecker()(
		upper.Alfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      "dead_pointer",
				Text:     "Referenced path does not exist",
				Path:     "Alfa",
				Type:     concern.Line,
				Line:     1,
				LineText: line,
			},
		},
		"",
		l,
	)
}

func TestPointersAbsolute(t *testing.T) {
	line := "Notes at `/Users/example/notes.md` today."
	l := pointerChecker()(
		upper.Alfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      "absolute_pointer",
				Text:     "Absolute path reference",
				Path:     "Alfa",
				Type:     concern.Line,
				Line:     1,
				LineText: line,
			},
		},
		"",
		l,
	)
}

func TestPointersRelativeExisting(t *testing.T) {
	l := pointerChecker("doc/ai/spec/naming.md")(
		"doc/ai/runbook/lint.md",
		strings.NewReader("See `../spec/naming.md` for the rules.\n"),
	)
	assertReport(t, "doc/ai/runbook/lint.md", false, nil, "", l)
}

func TestPointersRelativeDead(t *testing.T) {
	line := "See `../spec/ghost.md` for the rules."
	l := pointerChecker("doc/ai/spec/naming.md")(
		"doc/ai/runbook/lint.md",
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(
		t,
		"doc/ai/runbook/lint.md",
		true,
		[]*concern.Concern{
			{
				Key:      "dead_pointer",
				Text:     "Referenced path does not exist",
				Path:     "doc/ai/runbook/lint.md",
				Type:     concern.Line,
				Line:     1,
				LineText: line,
			},
		},
		"",
		l,
	)
}

func TestPointersSymbolIgnored(t *testing.T) {
	l := pointerChecker()(
		upper.Alfa,
		strings.NewReader(
			"Wraps `pkg/provision/salt.Client` and `pkg/system/run/New()`.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersExpansionExisting(t *testing.T) {
	l := pointerChecker("doc/ai/spec/naming.md", "doc/ai/spec/build.md")(
		upper.Alfa,
		strings.NewReader("Read `doc/ai/spec/{naming,build}.md` first.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersExpansionDead(t *testing.T) {
	line := "Read `doc/ai/spec/{naming,ghost}.md` first."
	l := pointerChecker("doc/ai/spec/naming.md")(
		upper.Alfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      "dead_pointer",
				Text:     "Referenced path does not exist",
				Path:     "Alfa",
				Type:     concern.Line,
				Line:     1,
				LineText: line,
			},
		},
		"",
		l,
	)
}

func TestPointersSiblingExisting(t *testing.T) {
	l := pointerChecker("../../github/soil/doc/ai/spec/naming.md")(
		upper.Alfa,
		strings.NewReader(
			"Shared specs at `../../github/soil/doc/ai/spec/naming.md`.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersSiblingDead(t *testing.T) {
	line := "Shared specs at `../github/soil/doc/ai/spec/naming.md`."
	l := pointerChecker()(
		upper.Alfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      "dead_pointer",
				Text:     "Referenced path does not exist",
				Path:     "Alfa",
				Type:     concern.Line,
				Line:     1,
				LineText: line,
			},
		},
		"",
		l,
	)
}

func TestPointersMultiple(t *testing.T) {
	line := "Read `doc/ai/spec/ghost.md` and `pkg/gone`."
	l := pointerChecker()(
		upper.Alfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      "dead_pointer",
				Text:     "Referenced path does not exist",
				Path:     "Alfa",
				Type:     concern.Line,
				Line:     1,
				LineText: line,
			},
			{
				Key:      "dead_pointer",
				Text:     "Referenced path does not exist",
				Path:     "Alfa",
				Type:     concern.Line,
				Line:     1,
				LineText: line,
			},
		},
		"",
		l,
	)
}

func TestPointersPluginRoot(t *testing.T) {
	l := pointerChecker("doc/ai/runbook/lint.md")(
		upper.Alfa,
		strings.NewReader(
			"Read `${CLAUDE_PLUGIN_ROOT}/doc/ai/runbook/lint.md`.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}
