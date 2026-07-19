package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestStrings(t *testing.T) {
	assert.Strings(t, []string{"a", "b"}, []string{"a", "b"})
}
