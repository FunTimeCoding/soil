package underscore

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestToDash(t *testing.T) {
	assert.String(t, "a-b-c", ToDash("a_b_c"))
}
