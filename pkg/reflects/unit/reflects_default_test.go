package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/reflects"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestDefault(t *testing.T) {
	assert.True(t, reflects.Default(false))
	assert.True(t, reflects.Default(0))
	assert.True(t, reflects.Default(""))
	assert.True(t, reflects.Default((*string)(nil)))
	assert.True(t, reflects.Default((*int)(nil)))
	assert.False(t, reflects.Default(nil))
	assert.False(t, reflects.Default(true))
	assert.False(t, reflects.Default(1))
	assert.False(t, reflects.Default(constant.UpperAlfa))
	assert.False(t, reflects.Default(new(constant.UpperAlfa)))
}
