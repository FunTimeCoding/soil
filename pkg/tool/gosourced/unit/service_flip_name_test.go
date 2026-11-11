package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/identifier"
	"testing"
)

func TestFlipNameUpperToLower(t *testing.T) {
	assert.String(t, "isGenerated", identifier.FlipName("IsGenerated"))
}

func TestFlipNameLowerToUpper(t *testing.T) {
	assert.String(t, "IsGenerated", identifier.FlipName("isGenerated"))
}

func TestFlipNameSingleCharacter(t *testing.T) {
	assert.String(t, "a", identifier.FlipName("A"))
	assert.String(t, "A", identifier.FlipName("a"))
}
