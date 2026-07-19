package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/integers"
	"testing"
)

func TestToStrings(t *testing.T) {
	assert.Any(t, []string{}, integers.ToStrings([]int{}))
	assert.Any(t, []string{"1"}, integers.ToStrings([]int{1}))
	assert.Any(t, []string{"2", "3"}, integers.ToStrings([]int{2, 3}))
}
