package system

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/join"
	"testing"
)

func TestFile(t *testing.T) {
	SaveFile(
		join.Absolute(WorkDirectory(), "test.txt"),
		"test content",
	)
	assert.String(
		t,
		"test content",
		ReadFile(WorkDirectory(), "test.txt"),
	)
	DeleteFile("test.txt")
}
