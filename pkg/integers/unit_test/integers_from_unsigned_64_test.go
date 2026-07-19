package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/integers"
	"testing"
)

func TestFromUnsigned64(t *testing.T) {
	assert.Integer(t, 0, integers.FromUnsigned64(0))
}
