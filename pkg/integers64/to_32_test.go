package integers64

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestTo32(t *testing.T) {
	assert.Integer32(t, 0, To32(0))
}
