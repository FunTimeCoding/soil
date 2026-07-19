package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"slices"
	"testing"
)

func TestSlices(t *testing.T) {
	assert.False(t, slices.Contains([]string{}, "a"))
	assert.False(t, slices.Contains([]string{}, ""))
}
