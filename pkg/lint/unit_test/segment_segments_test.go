package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/segment"
	"testing"
)

func TestSegmentsCamelCase(t *testing.T) {
	assert.Strings(t, []string{"dir", "name"}, segment.Segments("dirName"))
}

func TestSegmentsPascalCase(t *testing.T) {
	assert.Strings(
		t,
		[]string{"dir", "something"},
		segment.Segments("DirSomething"),
	)
}

func TestSegmentsSnakeCase(t *testing.T) {
	assert.Strings(
		t,
		[]string{"model", "context"},
		segment.Segments("model_context"),
	)
}

func TestSegmentsSingleWord(t *testing.T) {
	assert.Strings(t, []string{"directory"}, segment.Segments("directory"))
}

func TestSegmentsAllLower(t *testing.T) {
	assert.Strings(t, []string{"url"}, segment.Segments("url"))
}

func TestSegmentsMixed(t *testing.T) {
	assert.Strings(
		t,
		[]string{"output", "directory"},
		segment.Segments("OutputDirectory"),
	)
}

func TestSegmentsTrailingInitialism(t *testing.T) {
	assert.Strings(
		t,
		[]string{"marshal", "json"},
		segment.Segments("MarshalJSON"),
	)
}

func TestSegmentsLeadingInitialism(t *testing.T) {
	assert.Strings(t, []string{"url", "option"}, segment.Segments("URLOption"))
}

func TestSegmentsInteriorInitialism(t *testing.T) {
	assert.Strings(
		t,
		[]string{"ip", "configuration"},
		segment.Segments("IPConfiguration"),
	)
}

func TestSegmentsDigitBeforeWord(t *testing.T) {
	assert.Strings(
		t,
		[]string{"base64", "value"},
		segment.Segments("Base64Value"),
	)
}
