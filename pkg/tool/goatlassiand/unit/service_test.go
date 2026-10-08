package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/service"
	"testing"
)

func TestReplaceUniqueSingleMatch(t *testing.T) {
	result, e := service.ReplaceUnique("Hello world", "world", "there")
	assert.FatalOnError(t, e)
	assert.String(t, "Hello there", result)
}

func TestReplaceUniqueNotFound(t *testing.T) {
	_, e := service.ReplaceUnique("Hello world", "missing", "replacement")
	assert.Error(t, e)
	assert.StringContains(t, "old_text not found in page", e.Error())
}

func TestReplaceUniqueMultipleMatches(t *testing.T) {
	_, e := service.ReplaceUnique("foo bar foo baz foo", "foo", "qux")
	assert.Error(t, e)
	assert.StringContains(
		t,
		"old_text found 3 times, must be unique",
		e.Error(),
	)
}

func TestReplaceUniqueTwoMatches(t *testing.T) {
	_, e := service.ReplaceUnique("the cat and the dog", "the", "a")
	assert.Error(t, e)
	assert.StringContains(
		t,
		"old_text found 2 times, must be unique",
		e.Error(),
	)
}

func TestReplaceUniqueMatchAtStart(t *testing.T) {
	result, e := service.ReplaceUnique(
		"## Heading\n\nBody text",
		"## Heading",
		"## New Heading",
	)
	assert.FatalOnError(t, e)
	assert.String(t, "## New Heading\n\nBody text", result)
}

func TestReplaceUniqueMatchAtEnd(t *testing.T) {
	result, e := service.ReplaceUnique(
		"First paragraph\n\nLast line",
		"Last line",
		"Final line",
	)
	assert.FatalOnError(t, e)
	assert.String(t, "First paragraph\n\nFinal line", result)
}

func TestReplaceUniqueMultilineOldText(t *testing.T) {
	result, e := service.ReplaceUnique(
		"# Title\n\n- Alfa\n- Bravo\n\nEnd",
		"- Alfa\n- Bravo",
		"- Charlie\n- Delta\n- Echo",
	)
	assert.FatalOnError(t, e)
	assert.String(t, "# Title\n\n- Charlie\n- Delta\n- Echo\n\nEnd", result)
}

func TestReplaceUniqueWithEmpty(t *testing.T) {
	result, e := service.ReplaceUnique(
		"Keep this. Remove this. Keep that.",
		" Remove this.",
		"",
	)
	assert.FatalOnError(t, e)
	assert.String(t, "Keep this. Keep that.", result)
}

func TestReplaceUniqueSectionWithContent(t *testing.T) {
	result, e := service.ReplaceUnique(
		"Before\n\nAfter",
		"Before\n\nAfter",
		"Before\n\nNew middle section\n\nAfter",
	)
	assert.FatalOnError(t, e)
	assert.String(t, "Before\n\nNew middle section\n\nAfter", result)
}

func TestReplaceUniqueByContext(t *testing.T) {
	result, e := service.ReplaceUnique(
		"the cat sat on the mat",
		"the cat",
		"a dog",
	)
	assert.FatalOnError(t, e)
	assert.String(t, "a dog sat on the mat", result)
}
