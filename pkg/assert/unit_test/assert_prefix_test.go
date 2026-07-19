package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestPrefix(t *testing.T) {
	assert.Prefix(t, "", "ab")
	assert.Prefix(t, "a", "ab")
	assert.Prefix(t, "ab", "ab")
}
