package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestUnsignedToInteger(t *testing.T) {
	assert.Unsigned(t, 5, strings.ToUnsignedInteger(strings.Five, 0))
	assert.Unsigned(t, 5, strings.ToUnsignedInteger(" 5", 0))
}
