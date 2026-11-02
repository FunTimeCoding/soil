package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol"
	"testing"
)

func TestFlipNameUpperToLower(t *testing.T) {
	assert.String(t, "isGenerated", symbol.FlipName("IsGenerated"))
}

func TestFlipNameLowerToUpper(t *testing.T) {
	assert.String(t, "IsGenerated", symbol.FlipName("isGenerated"))
}

func TestFlipNameSingleCharacter(t *testing.T) {
	assert.String(t, "a", symbol.FlipName("A"))
	assert.String(t, "A", symbol.FlipName("a"))
}
