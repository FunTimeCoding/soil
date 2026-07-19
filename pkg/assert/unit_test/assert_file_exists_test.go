package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFileExists(t *testing.T) {
	assert.FileExists(t, "./assert_file_exists_test.go")
}
