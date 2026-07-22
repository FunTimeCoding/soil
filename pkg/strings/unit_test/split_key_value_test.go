package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
	"testing"
)

func TestAt(t *testing.T) {
	a, b := key_value.At("a@b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := key_value.At("c@d@e")
	assert.String(t, "c", c)
	assert.String(t, "d@e", d)
	f, g := key_value.At("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := key_value.At("h@")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}

func TestColon(t *testing.T) {
	a, b := key_value.Colon("a:b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := key_value.Colon("c:d:e")
	assert.String(t, "c", c)
	assert.String(t, "d:e", d)
	f, g := key_value.Colon("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := key_value.Colon("h:")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}

func TestComma(t *testing.T) {
	a, b := key_value.Comma("a,b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := key_value.Comma("c,d,e")
	assert.String(t, "c", c)
	assert.String(t, "d,e", d)
	f, g := key_value.Comma("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := key_value.Comma("h,")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}

func TestDash(t *testing.T) {
	a, b := key_value.Dash("a-b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := key_value.Dash("c-d-e")
	assert.String(t, "c", c)
	assert.String(t, "d-e", d)
	f, g := key_value.Dash("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := key_value.Dash("h-")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}

func TestDot(t *testing.T) {
	a, b := key_value.Dot("a.b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := key_value.Dot("c.d.e")
	assert.String(t, "c", c)
	assert.String(t, "d.e", d)
	f, g := key_value.Dot("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := key_value.Dot("h.")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}

func TestSplitKeyValueEquals(t *testing.T) {
	a, b := key_value.Equals("a=b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := key_value.Equals("c=d=e")
	assert.String(t, "c", c)
	assert.String(t, "d=e", d)
	f, g := key_value.Equals("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := key_value.Equals("h=")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}

func TestPipe(t *testing.T) {
	a, b := key_value.Pipe("a|b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := key_value.Pipe("c|d|e")
	assert.String(t, "c", c)
	assert.String(t, "d|e", d)
	f, g := key_value.Pipe("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := key_value.Pipe("h|")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}

func TestSemicolon(t *testing.T) {
	a, b := key_value.Semicolon("a;b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := key_value.Semicolon("c;d;e")
	assert.String(t, "c", c)
	assert.String(t, "d;e", d)
	f, g := key_value.Semicolon("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := key_value.Semicolon("h;")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}

func TestSlash(t *testing.T) {
	a, b := key_value.Slash("a/b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := key_value.Slash("c/d/e")
	assert.String(t, "c", c)
	assert.String(t, "d/e", d)
	f, g := key_value.Slash("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := key_value.Slash("h/")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}

func TestSpace(t *testing.T) {
	a, b := key_value.Space("a b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := key_value.Space("c d e")
	assert.String(t, "c", c)
	assert.String(t, "d e", d)
	f, g := key_value.Space("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := key_value.Space("h ")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}

func TestTab(t *testing.T) {
	a, b := key_value.Tab("a\tb")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := key_value.Tab("c\td\te")
	assert.String(t, "c", c)
	assert.String(t, "d\te", d)
	f, g := key_value.Tab("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := key_value.Tab("h\t")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}

func TestUnderscore(t *testing.T) {
	a, b := key_value.Underscore("a_b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := key_value.Underscore("c_d_e")
	assert.String(t, "c", c)
	assert.String(t, "d_e", d)
	f, g := key_value.Underscore("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := key_value.Underscore("h_")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}
