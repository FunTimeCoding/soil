package option

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.NotNil(t, Color)
	assert.NotNil(t, ExtendedColor)
}
