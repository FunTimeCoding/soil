package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestUnsigned(t *testing.T) {
	assert.Unsigned(t, 1, 1)
}
