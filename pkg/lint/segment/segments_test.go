package segment

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestSegmentsCamelCase(t *testing.T) {
	assert.Strings(t, []string{"dir", "name"}, Segments("dirName"))
}

func TestSegmentsPascalCase(t *testing.T) {
	assert.Strings(t, []string{"dir", "something"}, Segments("DirSomething"))
}

func TestSegmentsSnakeCase(t *testing.T) {
	assert.Strings(t, []string{"model", "context"}, Segments("model_context"))
}

func TestSegmentsSingleWord(t *testing.T) {
	assert.Strings(t, []string{"directory"}, Segments("directory"))
}

func TestSegmentsAllLower(t *testing.T) {
	assert.Strings(t, []string{"url"}, Segments("url"))
}

func TestSegmentsMixed(t *testing.T) {
	assert.Strings(
		t,
		[]string{"output", "directory"},
		Segments("OutputDirectory"),
	)
}
