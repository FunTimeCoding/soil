package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestWindowsToLinuxPath(t *testing.T) {
	assert.String(
		t,
		"/c/Users/example/Desktop/file.txt",
		system.WindowsToLinuxPath("C:\\Users\\example\\Desktop\\file.txt"),
	)
}
