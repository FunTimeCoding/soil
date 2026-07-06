package split

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestSplit(t *testing.T) {
	assert.Strings(t, []string{"a", "b"}, Colon("a:b"))
	assert.Strings(t, []string{"a", "b"}, Comma("a,b"))
	assert.Strings(t, []string{"a", "b"}, Dash("a-b"))
	assert.Strings(t, []string{"a", "b"}, Dot("a.b"))
	assert.Strings(t, []string{"a", "b"}, DoubleColon("a::b"))
	assert.Strings(t, []string{"a", "b"}, NewLine("a\nb"))
	assert.Strings(t, []string{"a", "b"}, Pipe("a|b"))
	assert.Strings(t, []string{"a", "b"}, Semicolon("a;b"))
	assert.Strings(t, []string{"a", "b"}, Slash("a/b"))
	assert.Strings(t, []string{"a", "b"}, Space("a b"))
	assert.Strings(t, []string{"a", "b"}, Underscore("a_b"))
}
