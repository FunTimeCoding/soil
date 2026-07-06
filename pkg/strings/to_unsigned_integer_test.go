package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestUnsignedToInteger(t *testing.T) {
	assert.Unsigned(t, 5, ToUnsignedInteger(Five, 0))
	assert.Unsigned(t, 5, ToUnsignedInteger(" 5", 0))
}
