package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/make_range"
	"testing"
)

func TestMakeRangeInteger(t *testing.T) {
	assert.Any(t, []int{0, 1}, make_range.Integer(0, 1))
}
