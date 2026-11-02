package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/reach"
	"testing"
)

func TestReachCommandTakesBothSidesOfAMove(t *testing.T) {
	assert.Strings(
		t,
		[]string{"/alfa/bravo.go", "/alfa/charlie/bravo.go"},
		reach.ResolveCommand("/alfa", "mv bravo.go charlie/bravo.go").Paths,
	)
}

func TestReachCommandTakesOnlyTheCopyDestination(t *testing.T) {
	assert.Strings(
		t,
		[]string{"/alfa/charlie.go"},
		reach.ResolveCommand("/alfa", "cp -f /tmp/bravo.go charlie.go").Paths,
	)
}

func TestReachCommandFollowsTheDirectoryChange(t *testing.T) {
	assert.Strings(
		t,
		[]string{"/alfa/bravo/charlie.go"},
		reach.ResolveCommand("/alfa", "cd bravo && rm -f charlie.go").Paths,
	)
}

func TestReachCommandTakesInPlaceSedOperandsAfterTheScript(t *testing.T) {
	assert.Strings(
		t,
		[]string{"/alfa/bravo.go", "/alfa/charlie.md"},
		reach.ResolveCommand(
			"/alfa",
			"gsed -i 's/delta/echo/' bravo.go charlie.md",
		).Paths,
	)
}

func TestReachCommandLeavesAReadingSedAlone(t *testing.T) {
	assert.Count(
		t,
		0,
		reach.ResolveCommand("/alfa", "sed -n 1,5p bravo.go").Paths,
	)
}

func TestReachCommandTakesALiteralRedirectButNotTheNullDevice(t *testing.T) {
	assert.Strings(
		t,
		[]string{"/alfa/bravo.md"},
		reach.ResolveCommand(
			"/alfa",
			"echo charlie > bravo.md 2>&1; echo delta > /dev/null",
		).Paths,
	)
}

func TestReachCommandLeavesAVariableOperandUnresolved(t *testing.T) {
	w := reach.ResolveCommand("/alfa", "for f in *.go; do rm $f; done")
	assert.Count(t, 0, w.Paths)
	assert.True(t, w.Unresolved)
}

func TestReachCommandNamesAFixerRunButNotItsDiff(t *testing.T) {
	assert.True(t, reach.ResolveCommand("/alfa", "gofix ./...").Fixer)
	assert.False(t, reach.ResolveCommand("/alfa", "gofix --diff ./...").Fixer)
}
