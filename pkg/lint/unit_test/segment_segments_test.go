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
