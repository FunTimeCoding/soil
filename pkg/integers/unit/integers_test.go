package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/integers"
	"testing"
)

func TestContains(t *testing.T) {
	assert.True(t, integers.Contains([]int{0}, 0))
	assert.False(t, integers.Contains([]int{0}, 1))
}

func TestFromUnsigned64(t *testing.T) {
	assert.Integer(t, 0, integers.FromUnsigned64(0))
}

func TestNextFreeNumber(t *testing.T) {
	assert.Integer(t, 0, integers.NextFree(0, []int{}))
	assert.Integer(t, 1, integers.NextFree(0, []int{0}))
	assert.Integer(t, 1, integers.NextFree(0, []int{0, 2}))
}

func TestRemoveFromList(t *testing.T) {
	assert.Any(
		t,
		[]int{2, 3},
		integers.RemoveFromList([]int{1, 1, 2, 3}, []int{1}),
	)
	assert.Any(
		t,
		[]int{3},
		integers.RemoveFromList([]int{1, 1, 2, 3}, []int{1, 2}),
	)
}

func TestTo32(t *testing.T) {
	assert.Integer(t, 0, integers.To32(0))
}

func TestToString(t *testing.T) {
	assert.String(t, "1", integers.ToString(1))
}

func TestToStrings(t *testing.T) {
	assert.Any(t, []string{}, integers.ToStrings([]int{}))
	assert.Any(t, []string{"1"}, integers.ToStrings([]int{1}))
	assert.Any(t, []string{"2", "3"}, integers.ToStrings([]int{2, 3}))
}
