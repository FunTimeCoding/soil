package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"testing"
)

func TestFlipNameUpperToLower(t *testing.T) {
	assert.String(t, "isGenerated", service.FlipName("IsGenerated"))
}

func TestFlipNameLowerToUpper(t *testing.T) {
	assert.String(t, "IsGenerated", service.FlipName("isGenerated"))
}

func TestFlipNameSingleCharacter(t *testing.T) {
	assert.String(t, "a", service.FlipName("A"))
	assert.String(t, "A", service.FlipName("a"))
}
