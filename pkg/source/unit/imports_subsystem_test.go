package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"testing"
)

func TestSubsystemIsTheParentDirectory(t *testing.T) {
	assert.String(
		t,
		"web",
		imports.Subsystem(
			"example.org/alfa/pkg/web/constant",
			"example.org/alfa",
		),
	)
}

func TestSubsystemAtTheLayoutRootIsTheModuleName(t *testing.T) {
	assert.String(
		t,
		"alfa",
		imports.Subsystem("example.org/alfa/pkg/constant", "example.org/alfa"),
	)
}

func TestSubsystemClimbsPastABannedName(t *testing.T) {
	assert.String(
		t,
		"machinery",
		imports.Subsystem(
			"example.org/machinery/pkg/api/errors",
			"example.org/machinery",
		),
	)
}

func TestSubsystemClimbsPastANameThatIsNoIdentifier(t *testing.T) {
	assert.String(
		t,
		"bravo",
		imports.Subsystem(
			"example.org/bravo/pkg/util-stuff/constant",
			"example.org/bravo",
		),
	)
}

func TestSubsystemOfAModuleRootIsEmpty(t *testing.T) {
	assert.String(
		t,
		"",
		imports.Subsystem("example.org/charlie-go", "example.org/charlie-go"),
	)
}

func TestSubsystemNeverClimbsAboveTheModule(t *testing.T) {
	assert.String(
		t,
		"",
		imports.Subsystem(
			"example.org/delta-sdk/alert",
			"example.org/delta-sdk",
		),
	)
}

func TestSubsystemOfTheStandardLibrary(t *testing.T) {
	assert.String(t, "crypto", imports.Subsystem("crypto/rand", ""))
}
