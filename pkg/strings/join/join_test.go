package join

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestJoin(t *testing.T) {
	s := []string{"a", "b"}
	assert.String(t, "a:b", Colon(s...))
	assert.String(t, "a,b", Comma(s))
	assert.String(t, "a, b", CommaSpace(s))
	assert.String(t, "a-b", Dash(s))
	assert.String(t, "a::b", DoubleColon(s))
	assert.String(t, "a=b", Equals(s))
	assert.String(t, "a\nb", NewLine(s))
	assert.String(t, "a|b", Pipe(s))
	assert.String(t, "a b", Space(s...))
	assert.String(t, "a_b", Underscore(s))
	assert.String(t, "a.b", Dot(s))
	assert.String(t, "a/b", Slash(s))
	assert.String(t, "ab", Empty(s...))
}
