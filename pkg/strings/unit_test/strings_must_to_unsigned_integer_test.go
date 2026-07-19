package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestMustToUnsignedInteger(t *testing.T) {
	assert.Unsigned(t, 1, strings.MustToUnsignedInteger(strings.One))
	assert.Unsigned(t, 1, strings.MustToUnsignedInteger(" 1"))
}
