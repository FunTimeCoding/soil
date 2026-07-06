package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestToInteger(t *testing.T) {
	assert.Integer(t, 5, ToInteger(Five, 0))
	assert.Integer(t, 5, ToInteger(" 5", 0))
}
