package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestMustToIntegers64(t *testing.T) {
	assert.Any(t, []int64{}, strings.MustToIntegers64([]string{}))
	assert.Any(t, []int64{5}, strings.MustToIntegers64([]string{"5"}))
	assert.Any(t, []int64{1, 2}, strings.MustToIntegers64([]string{"1", "2"}))
}
