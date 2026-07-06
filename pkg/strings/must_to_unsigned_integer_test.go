package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestMustToUnsignedInteger(t *testing.T) {
	assert.Unsigned(t, 1, MustToUnsignedInteger(One))
	assert.Unsigned(t, 1, MustToUnsignedInteger(" 1"))
}
