package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestContainsLower(t *testing.T) {
	assert.True(t, strings.ContainsLower("TestFixture", "fixture"))
	assert.False(t, strings.ContainsLower("TestFixture", "other"))
}

func TestCountRune(t *testing.T) {
	assert.Integer(t, 0, strings.CountRune("a", 'b'))
	assert.Integer(t, 1, strings.CountRune("a", 'a'))
	assert.Integer(t, 2, strings.CountRune("aa", 'a'))
}

func TestElementAfterSkip(t *testing.T) {
	assert.String(
		t,
		"b",
		strings.ElementAfterSkip([]string{"a", "b", "a", "c"}, "a", 0),
	)
	assert.String(
		t,
		"c",
		strings.ElementAfterSkip([]string{"a", "b", "a", "c"}, "a", 1),
	)
}

func TestElementAfter(t *testing.T) {
	assert.String(
		t,
		"b",
		strings.ElementAfter([]string{"a", "b"}, "a"),
	)
}

func TestHas(t *testing.T) {
	assert.True(t, strings.Has("abc", "a", "c"))
	assert.False(t, strings.Has("abc", "a", "d"))
	assert.False(t, strings.Has("abc", "x", "c"))
	assert.False(t, strings.Has("abc", "x", "d"))
}

func TestHasWord(t *testing.T) {
	assert.True(t, strings.HasWord("alfa bravo", "bravo"))
	assert.False(t, strings.HasWord("alfa bravo", "bro"))
	assert.False(t, strings.HasWord("alfa bravo", "bra"))
	assert.False(t, strings.HasWord("alfa bravo", "charlie"))
}

func TestIndexOfSkip(t *testing.T) {
	assert.Integer(
		t,
		1,
		strings.IndexOfSkip(
			upper.Bravo,
			[]string{upper.Alfa, upper.Bravo, upper.Charlie, upper.Bravo},
			0,
		),
	)
	assert.Integer(
		t,
		3,
		strings.IndexOfSkip(
			upper.Bravo,
			[]string{upper.Alfa, upper.Bravo, upper.Charlie, upper.Bravo},
			1,
		),
	)
}

func TestIndexOf(t *testing.T) {
	assert.Integer(
		t,
		1,
		strings.IndexOf(
			upper.Bravo,
			[]string{upper.Alfa, upper.Bravo, upper.Charlie},
		),
	)
}

func TestIntensity(t *testing.T) {
	v := []string{upper.Alfa, upper.Bravo, upper.Charlie}
	assert.String(t, "Alfa", strings.Intensity(v, 0))
	assert.String(t, "Alfa", strings.Intensity(v, 0.33))
	assert.String(t, "Bravo", strings.Intensity(v, 0.34))
	assert.String(t, "Bravo", strings.Intensity(v, 0.66))
	assert.String(t, "Charlie", strings.Intensity(v, 0.67))
	assert.String(t, "Charlie", strings.Intensity(v, 1))
}
