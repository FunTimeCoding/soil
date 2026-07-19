package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/reflects"
	"testing"
)

func TestNilPointer(t *testing.T) {
	assert.True(t, reflects.NilPointer((*int)(nil)))
	assert.False(t, reflects.NilPointer(42))
}
