package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"testing"
)

func TestMockAliasCarriesItsPackage(t *testing.T) {
	assert.String(
		t,
		"bravoMock",
		imports.MockAlias(
			"example.org/alfa/pkg/bravo/mock_client",
			"example.org/alfa",
		),
	)
}

func TestMockAliasCamelsSnakeOwner(t *testing.T) {
	assert.String(
		t,
		"bravoCharlieMock",
		imports.MockAlias(
			"example.org/alfa/pkg/bravo_charlie/mock_client",
			"example.org/alfa",
		),
	)
}

func TestMockAliasClimbsPastClientToItsOwner(t *testing.T) {
	assert.String(
		t,
		"goalfaMock",
		imports.MockAlias(
			"example.org/alfa/pkg/tool/goalfa/client/mock_client",
			"example.org/alfa",
		),
	)
}

func TestMockAliasClimbsPastUnitToItsOwner(t *testing.T) {
	assert.String(
		t,
		"goalfaMock",
		imports.MockAlias(
			"example.org/alfa/pkg/tool/goalfa/unit/mock_client",
			"example.org/alfa",
		),
	)
}
