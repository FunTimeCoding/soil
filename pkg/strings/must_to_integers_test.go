package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestMustToIntegers(t *testing.T) {
	assert.Any(t, []int{}, MustToIntegers([]string{}))
	assert.Any(t, []int{5}, MustToIntegers([]string{"5"}))
	assert.Any(t, []int{1, 2}, MustToIntegers([]string{"1", "2"}))
}
