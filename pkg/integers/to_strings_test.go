package integers

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestToStrings(t *testing.T) {
	assert.Any(t, []string{}, ToStrings([]int{}))
	assert.Any(t, []string{"1"}, ToStrings([]int{1}))
	assert.Any(t, []string{"2", "3"}, ToStrings([]int{2, 3}))
}
