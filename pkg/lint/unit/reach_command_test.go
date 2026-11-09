package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/reach"
	"testing"
)

func TestReachCommandTakesBothSidesOfMove(t *testing.T) {
	assert.Strings(
		t,
		[]string{"/alfa/bravo.go", "/alfa/charlie/bravo.go"},
		reach.ResolveCommand("/alfa", "mv bravo.go charlie/bravo.go").Paths,
	)
}

func TestReachCommandTakesOnlyCopyDestination(t *testing.T) {
	assert.Strings(
		t,
		[]string{"/alfa/charlie.go"},
		reach.ResolveCommand("/alfa", "cp -f /tmp/bravo.go charlie.go").Paths,
	)
}

func TestReachCommandFollowsDirectoryChange(t *testing.T) {
	assert.Strings(
		t,
		[]string{"/alfa/bravo/charlie.go"},
		reach.ResolveCommand("/alfa", "cd bravo && rm -f charlie.go").Paths,
	)
}

func TestReachCommandTakesInPlaceSedOperandsAfterScript(t *testing.T) {
	assert.Strings(
		t,
		[]string{"/alfa/bravo.go", "/alfa/charlie.md"},
		reach.ResolveCommand(
			"/alfa",
			"gsed -i 's/delta/echo/' bravo.go charlie.md",
		).Paths,
	)
}

func TestReachCommandLeavesReadingSedAlone(t *testing.T) {
	assert.Count(
		t,
		0,
		reach.ResolveCommand("/alfa", "sed -n 1,5p bravo.go").Paths,
	)
}

func TestReachCommandTakesLiteralRedirectButNotNullDevice(t *testing.T) {
	assert.Strings(
		t,
		[]string{"/alfa/bravo.md"},
		reach.ResolveCommand(
			"/alfa",
			"echo charlie > bravo.md 2>&1; echo delta > /dev/null",
		).Paths,
	)
}

func TestReachCommandLeavesVariableOperandUnresolved(t *testing.T) {
	w := reach.ResolveCommand("/alfa", "for f in *.go; do rm $f; done")
	assert.Count(t, 0, w.Paths)
	assert.True(t, w.Unresolved)
}

func TestReachCommandNamesFixerRunButNotItsDiff(t *testing.T) {
	assert.True(t, reach.ResolveCommand("/alfa", "gofix ./...").Fixer)
	assert.False(t, reach.ResolveCommand("/alfa", "gofix --diff ./...").Fixer)
}
