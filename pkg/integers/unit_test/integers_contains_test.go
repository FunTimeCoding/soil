package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/integers"
	"testing"
)

func TestContains(t *testing.T) {
	assert.True(t, integers.Contains([]int{0}, 0))
	assert.False(t, integers.Contains([]int{0}, 1))
}
