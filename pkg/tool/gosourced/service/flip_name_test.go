package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFlipNameUpperToLower(t *testing.T) {
	assert.String(t, "isGenerated", flipName("IsGenerated"))
}

func TestFlipNameLowerToUpper(t *testing.T) {
	assert.String(t, "IsGenerated", flipName("isGenerated"))
}

func TestFlipNameSingleCharacter(t *testing.T) {
	assert.String(t, "a", flipName("A"))
	assert.String(t, "A", flipName("a"))
}
