package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/integers64"
	"testing"
)

func TestToUnsigned32(t *testing.T) {
	assert.Unsigned32(t, 0, integers64.ToUnsigned32(0))
}
