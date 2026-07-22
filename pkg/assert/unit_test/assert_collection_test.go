package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestContains(t *testing.T) {
	assert.Contains(t, []string{}, []string{"a", "b"})
	assert.Contains(t, []string{"a"}, []string{"a", "b"})
	assert.Contains(t, []string{"a", "b"}, []string{"a", "b"})
}

func TestCount(t *testing.T) {
	assert.Count(t, 0, []string{})
	assert.Count(t, 1, []string{"1"})
	assert.Count(t, 2, []string{"1", "2"})
}

func TestIntegers(t *testing.T) {
	assert.Integers(t, []int{0, 1}, []int{0, 1})
}

func TestMapHasKey(t *testing.T) {
	m := map[string]any{"height": 180}
	assert.MapHasKey(t, m, "height")
}

func TestMapNotHasKey(t *testing.T) {
	m := map[string]any{"height": 180}
	assert.MapNotHasKey(t, m, "weight")
}

func TestMapValue(t *testing.T) {
	m := map[string]any{"height": 180, "name": "Fern"}
	assert.MapValue(t, 180, m, "height")
	assert.MapValue(t, "Fern", m, "name")
}

func TestPrefix(t *testing.T) {
	assert.Prefix(t, "", "ab")
	assert.Prefix(t, "a", "ab")
	assert.Prefix(t, "ab", "ab")
}

func TestStringContains(t *testing.T) {
	assert.StringContains(t, "friend", "hello friend")
}

func TestStringNotContains(t *testing.T) {
	assert.StringNotContains(t, "enemy", "hello friend")
}

func TestStrings(t *testing.T) {
	assert.Strings(t, []string{"a", "b"}, []string{"a", "b"})
}

func TestSuffix(t *testing.T) {
	assert.Suffix(t, "", "ab")
	assert.Suffix(t, "b", "ab")
	assert.Suffix(t, "ab", "ab")
}
