package integers

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestRemoveFromList(t *testing.T) {
	assert.Any(
		t,
		[]int{2, 3},
		RemoveFromList([]int{1, 1, 2, 3}, []int{1}),
	)
	assert.Any(
		t,
		[]int{3},
		RemoveFromList([]int{1, 1, 2, 3}, []int{1, 2}),
	)
}
