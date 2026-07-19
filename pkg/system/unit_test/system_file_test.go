package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/join"
	"testing"
)

func TestFile(t *testing.T) {
	system.SaveFile(
		join.Absolute(system.WorkDirectory(), "test.txt"),
		"test content",
	)
	assert.String(
		t,
		"test content",
		system.ReadFile(system.WorkDirectory(), "test.txt"),
	)
	system.DeleteFile("test.txt")
}
