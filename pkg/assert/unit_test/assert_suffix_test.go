package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestSuffix(t *testing.T) {
	assert.Suffix(t, "", "ab")
	assert.Suffix(t, "b", "ab")
	assert.Suffix(t, "ab", "ab")
}
