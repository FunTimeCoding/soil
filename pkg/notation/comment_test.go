package notation

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestComment(t *testing.T) {
	assert.String(t, "a:\n\"b\"", Comment("a", "b"))
}
