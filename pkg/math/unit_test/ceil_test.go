package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/ceil"
	"testing"
)

func TestCeilInteger(t *testing.T) {
	assert.Integer(t, 1, ceil.Integer(0.9))
}
