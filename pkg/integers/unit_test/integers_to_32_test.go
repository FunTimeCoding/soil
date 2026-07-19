package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/integers"
	"testing"
)

func TestTo32(t *testing.T) {
	assert.Integer32(t, 0, integers.To32(0))
}
