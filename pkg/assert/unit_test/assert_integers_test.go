package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestIntegers(t *testing.T) {
	assert.Integers(t, []int{0, 1}, []int{0, 1})
}
