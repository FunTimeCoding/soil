package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/integers"
	"testing"
)

func TestNextFreeNumber(t *testing.T) {
	assert.Integer(t, 0, integers.NextFree(0, []int{}))
	assert.Integer(t, 1, integers.NextFree(0, []int{0}))
	assert.Integer(t, 1, integers.NextFree(0, []int{0, 2}))
}
