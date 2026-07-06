package lint

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestMissingTestDirectories(t *testing.T) {
	paths := []string{
		"pkg/foo/foo.go",
		"pkg/bar/bar.go",
		"pkg/bar/bar_test.go",
	}
	result := missingTestDirectories(paths, nil)
	assert.Integer(t, 1, len(result))
	_, okay := result["pkg/foo"]
	assert.Boolean(t, true, okay)
}

func TestMissingTestDirectoriesNone(t *testing.T) {
	paths := []string{
		"pkg/foo/foo.go",
		"pkg/foo/foo_test.go",
	}
	result := missingTestDirectories(paths, nil)
	assert.Integer(t, 0, len(result))
}

func TestMissingTestDirectoriesGenerated(t *testing.T) {
	paths := []string{
		"pkg/foo/foo.go",
	}
	generatedPaths := []string{
		"pkg/baz/generated.go",
	}
	result := missingTestDirectories(paths, generatedPaths)
	assert.Integer(t, 2, len(result))
	_, okay := result["pkg/foo"]
	assert.Boolean(t, true, okay)
	_, okay = result["pkg/baz"]
	assert.Boolean(t, true, okay)
}
