package constant

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "str", Str)
	assert.String(t, "int", Int)
	assert.String(t, "float", Float)
}
