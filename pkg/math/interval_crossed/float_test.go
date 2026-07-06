package interval_crossed

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFloat(t *testing.T) {
	assert.False(t, Float(10, 0, 1))
	assert.False(t, Float(10, 1, 9))
	assert.True(t, Float(10, 9, 10))
	assert.True(t, Float(10, 9, 11))
	assert.False(t, Float(10, 10, 11))
}
