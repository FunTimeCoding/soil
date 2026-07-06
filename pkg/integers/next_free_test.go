package integers

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestNextFreeNumber(t *testing.T) {
	assert.Integer(t, 0, NextFree(0, []int{}))
	assert.Integer(t, 1, NextFree(0, []int{0}))
	assert.Integer(t, 1, NextFree(0, []int{0, 2}))
}
