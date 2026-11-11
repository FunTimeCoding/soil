package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"testing"
)

func TestConstantAliasCarriesItsDomain(t *testing.T) {
	assert.String(
		t,
		"bravoConstant",
		imports.ConstantAlias(
			"example.org/alfa/pkg/bravo/constant",
			"example.org/alfa",
		),
	)
}

func TestConstantAliasCamelsSnakeDomain(t *testing.T) {
	assert.String(
		t,
		"bravoCharlieConstant",
		imports.ConstantAlias(
			"example.org/alfa/pkg/bravo_charlie/constant",
			"example.org/alfa",
		),
	)
}
