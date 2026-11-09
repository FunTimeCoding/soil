package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/reference"
	"testing"
)

func TestLivePathInBackticksPasses(t *testing.T) {
	assert.Count(t, 0, check("See `doc/real.md` first."))
}

func TestDeadPathInBackticksIsReported(t *testing.T) {
	assert.Any(
		t,
		[]*reference.Finding{
			reference.NewFinding(
				"doc/gone.md",
				"Referenced path does not exist",
			),
		},
		check("See `doc/gone.md` first."),
	)
}

func TestPathInProseAsksForBackticks(t *testing.T) {
	assert.Any(
		t,
		[]*reference.Finding{
			reference.NewFinding(
				"doc/real.md",
				"Path in prose - wrap it in backticks so it can be checked",
			),
		},
		check("Read doc/real.md before editing."),
	)
}

func TestProsePunctuationIsNotPartOfPath(t *testing.T) {
	assert.Any(
		t,
		[]*reference.Finding{
			reference.NewFinding(
				"doc/real.md",
				"Path in prose - wrap it in backticks so it can be checked",
			),
			reference.NewFinding(
				"pkg/tool",
				"Path in prose - wrap it in backticks so it can be checked",
			),
		},
		check("**(see doc/real.md)** and pkg/tool's helper"),
	)
}

func TestSiblingPathInProseKeepsItsParentPrefix(t *testing.T) {
	assert.Any(
		t,
		[]*reference.Finding{
			reference.NewFinding(
				"../soil/doc/real.md",
				"Path in prose - wrap it in backticks so it can be checked",
			),
		},
		check("(the economy is in ../soil/doc/real.md)."),
	)
}

func TestLocatorInProseIsNotPath(t *testing.T) {
	assert.Count(t, 0, check("Docs at https://example.com/doc/real.md."))
}

func TestCitedMemoryThatExistsPasses(t *testing.T) {
	assert.Count(
		t,
		0,
		check("Pairs with `memory://default/terse communication`."),
	)
}

func TestCitedMemoryThatIsMissingIsReported(t *testing.T) {
	assert.Any(
		t,
		[]*reference.Finding{
			reference.NewFinding(
				"memory://default/gone memory",
				"Referenced memory does not exist",
			),
		},
		check("Pairs with `memory://default/gone memory`."),
	)
}

func TestLiveHeadingPasses(t *testing.T) {
	assert.Count(t, 0, check("See `doc/real.md#usage` first."))
}

func TestDeadHeadingIsReportedWithNearest(t *testing.T) {
	assert.Any(
		t,
		[]*reference.Finding{
			reference.NewFinding(
				"doc/real.md#usages",
				"Referenced heading does not exist - nearest: #usage (\"Usage\"), #real (\"Real\")",
			),
		},
		check("See `doc/real.md#usages` first."),
	)
}

func TestHeadingOnDeadPathReportsPath(t *testing.T) {
	assert.Any(
		t,
		[]*reference.Finding{
			reference.NewFinding(
				"doc/gone.md#usage",
				"Referenced path does not exist",
			),
		},
		check("See `doc/gone.md#usage` first."),
	)
}

func TestBareNameUnderBaseResolves(t *testing.T) {
	assert.Count(t, 0, checkUnder("See `real.md` first.", "doc"))
}

func TestDeadBareNameUnderBaseIsReported(t *testing.T) {
	assert.Any(
		t,
		[]*reference.Finding{
			reference.NewFinding("gone.md", "Referenced path does not exist"),
		},
		checkUnder("See `gone.md` first.", "doc"),
	)
}

func TestBareNameTriesEveryBase(t *testing.T) {
	assert.Count(t, 0, checkUnder("See `real.go` and `real.md`.", "doc", "pkg"))
}

func TestHeadingUnderBaseIsChecked(t *testing.T) {
	assert.Any(
		t,
		[]*reference.Finding{
			reference.NewFinding(
				"real.md#usages",
				"Referenced heading does not exist - nearest: #usage (\"Usage\"), #real (\"Real\")",
			),
		},
		checkUnder("See `real.md#usages` first.", "doc"),
	)
}

func TestDeadBaseIsReported(t *testing.T) {
	assert.Any(
		t,
		[]*reference.Finding{
			reference.NewFinding("doc/gone", "Referenced path does not exist"),
		},
		checkUnder("Nothing cited.", "doc/gone"),
	)
}

func TestIdentifierUnderBaseIsNotPath(t *testing.T) {
	assert.Count(
		t,
		0,
		checkUnder("Call `w.poll`, `response.Fail` and `v.Detail`.", "doc"),
	)
}

func TestBareNameWithoutBaseIsNotChecked(t *testing.T) {
	assert.Count(t, 0, check("See `gone.md` first."))
}

func TestFragmentOnNonMarkdownTargetIsReported(t *testing.T) {
	assert.Any(
		t,
		[]*reference.Finding{
			reference.NewFinding(
				"pkg/real.go#check",
				"Fragment on a non-markdown target - only markdown headings can be referenced",
			),
		},
		check("See `pkg/real.go#check` first."),
	)
}
