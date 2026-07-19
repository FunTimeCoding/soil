package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"testing"
)

func TestJoin(t *testing.T) {
	s := []string{"a", "b"}
	assert.String(t, "a:b", join.Colon(s...))
	assert.String(t, "a,b", join.Comma(s))
	assert.String(t, "a, b", join.CommaSpace(s))
	assert.String(t, "a-b", join.Dash(s))
	assert.String(t, "a::b", join.DoubleColon(s))
	assert.String(t, "a=b", join.Equals(s))
	assert.String(t, "a\nb", join.NewLine(s))
	assert.String(t, "a|b", join.Pipe(s))
	assert.String(t, "a b", join.Space(s...))
	assert.String(t, "a_b", join.Underscore(s))
	assert.String(t, "a.b", join.Dot(s))
	assert.String(t, "a/b", join.Slash(s))
	assert.String(t, "ab", join.Empty(s...))
}
