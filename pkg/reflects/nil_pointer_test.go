package reflects

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestNilPointer(t *testing.T) {
	assert.True(t, NilPointer((*int)(nil)))
	assert.False(t, NilPointer(42))
}
