package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"testing"
)

func TestSplit(t *testing.T) {
	assert.Strings(t, []string{"a", "b"}, split.Colon("a:b"))
	assert.Strings(t, []string{"a", "b"}, split.Comma("a,b"))
	assert.Strings(t, []string{"a", "b"}, split.Dash("a-b"))
	assert.Strings(t, []string{"a", "b"}, split.Dot("a.b"))
	assert.Strings(t, []string{"a", "b"}, split.DoubleColon("a::b"))
	assert.Strings(t, []string{"a", "b"}, split.NewLine("a\nb"))
	assert.Strings(t, []string{"a", "b"}, split.Pipe("a|b"))
	assert.Strings(t, []string{"a", "b"}, split.Semicolon("a;b"))
	assert.Strings(t, []string{"a", "b"}, split.Slash("a/b"))
	assert.Strings(t, []string{"a", "b"}, split.Space("a b"))
	assert.Strings(t, []string{"a", "b"}, split.Underscore("a_b"))
}
