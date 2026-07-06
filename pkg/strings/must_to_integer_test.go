package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestMustToInteger(t *testing.T) {
	assert.Integer(t, 1, MustToInteger(One))
	assert.Integer(t, 1, MustToInteger(" 1"))
}
