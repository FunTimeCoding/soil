package reflects

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestDefault(t *testing.T) {
	assert.True(t, Default(false))
	assert.True(t, Default(0))
	assert.True(t, Default(""))
	assert.True(t, Default((*string)(nil)))
	assert.True(t, Default((*int)(nil)))
	assert.False(t, Default(nil))
	assert.False(t, Default(true))
	assert.False(t, Default(1))
	assert.False(t, Default(upper.Alfa))
	assert.False(t, Default(new(upper.Alfa)))
}
