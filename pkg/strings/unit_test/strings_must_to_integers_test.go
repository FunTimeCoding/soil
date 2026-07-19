package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestMustToIntegers(t *testing.T) {
	assert.Any(t, []int{}, strings.MustToIntegers([]string{}))
	assert.Any(t, []int{5}, strings.MustToIntegers([]string{"5"}))
	assert.Any(t, []int{1, 2}, strings.MustToIntegers([]string{"1", "2"}))
}
