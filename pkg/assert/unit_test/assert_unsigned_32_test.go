package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestUnsigned32(t *testing.T) {
	assert.Unsigned32(t, 1, 1)
}
