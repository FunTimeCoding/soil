package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/reference"
	"testing"
)

func TestALivePathInBackticksPasses(t *testing.T) {
	assert.Count(t, 0, check("See `doc/real.md` first."))
}

func TestADeadPathInBackticksIsReported(t *testing.T) {
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

func TestAPathInProseAsksForBackticks(t *testing.T) {
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

func TestProsePunctuationIsNotPartOfThePath(t *testing.T) {
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

func TestASiblingPathInProseKeepsItsParentPrefix(t *testing.T) {
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

func TestALocatorInProseIsNotAPath(t *testing.T) {
	assert.Count(t, 0, check("Docs at https://example.com/doc/real.md."))
}

func TestACitedMemoryThatExistsPasses(t *testing.T) {
	assert.Count(
		t,
		0,
		check("Pairs with `memory://default/terse communication`."),
	)
}

func TestACitedMemoryThatIsMissingIsReported(t *testing.T) {
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
