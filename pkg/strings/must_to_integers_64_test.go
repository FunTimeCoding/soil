package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestMustToIntegers64(t *testing.T) {
	assert.Any(t, []int64{}, MustToIntegers64([]string{}))
	assert.Any(t, []int64{5}, MustToIntegers64([]string{"5"}))
	assert.Any(t, []int64{1, 2}, MustToIntegers64([]string{"1", "2"}))
}
