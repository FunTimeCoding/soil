package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestMustToFloat(t *testing.T) {
	assert.Float(t, 1, MustToFloat("1.0"))
	assert.Float(t, 1, MustToFloat(" 1.0"))
}
