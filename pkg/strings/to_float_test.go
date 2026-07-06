package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestToFloat(t *testing.T) {
	assert.Float(t, 0.5, ToFloat("0.5", 0))
	assert.Float(t, 0.5, ToFloat(" 0.5", 0))
}
