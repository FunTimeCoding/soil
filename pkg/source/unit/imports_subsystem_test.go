package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"testing"
)

func TestSubsystemIsParentDirectory(t *testing.T) {
	assert.String(
		t,
		"web",
		imports.Subsystem(
			"example.org/alfa/pkg/web/constant",
			"example.org/alfa",
		),
	)
}

func TestSubsystemAtLayoutRootIsModuleName(t *testing.T) {
	assert.String(
		t,
		"alfa",
		imports.Subsystem("example.org/alfa/pkg/constant", "example.org/alfa"),
	)
}

func TestSubsystemClimbsPastBannedName(t *testing.T) {
	assert.String(
		t,
		"machinery",
		imports.Subsystem(
			"example.org/machinery/pkg/api/errors",
			"example.org/machinery",
		),
	)
}

func TestSubsystemClimbsPastNameThatIsNoIdentifier(t *testing.T) {
	assert.String(
		t,
		"bravo",
		imports.Subsystem(
			"example.org/bravo/pkg/util-stuff/constant",
			"example.org/bravo",
		),
	)
}

func TestSubsystemOfModuleRootIsEmpty(t *testing.T) {
	assert.String(
		t,
		"",
		imports.Subsystem("example.org/charlie-go", "example.org/charlie-go"),
	)
}

func TestSubsystemNeverClimbsAboveModule(t *testing.T) {
	assert.String(
		t,
		"",
		imports.Subsystem(
			"example.org/delta-sdk/alert",
			"example.org/delta-sdk",
		),
	)
}

func TestSubsystemOfStandardLibrary(t *testing.T) {
	assert.String(t, "crypto", imports.Subsystem("crypto/rand", ""))
}
