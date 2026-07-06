package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestMustToInteger64(t *testing.T) {
	assert.Integer64(t, 1, MustToInteger64(One))
	assert.Integer64(t, 1, MustToInteger64(" 1"))
}
