package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"testing"
)

func TestComment(t *testing.T) {
	assert.String(t, "a:\n\"b\"", notation.Comment("a", "b"))
}
