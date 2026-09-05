package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/join"
	"testing"
)

func TestReplaceFile(t *testing.T) {
	d := t.TempDir()
	source := join.Absolute(d, "source")
	destination := join.Absolute(d, "destination")
	system.SaveFile(source, "new")
	system.Executable(source)
	system.SaveFile(destination, "old")
	system.ReplaceFile(source, destination)
	assert.String(t, "new", system.ReadFile(d, "destination"))
	assert.True(t, system.IsExecutable(destination))
}

func TestReplaceFileCreatesDestination(t *testing.T) {
	d := t.TempDir()
	source := join.Absolute(d, "source")
	destination := join.Absolute(d, "destination")
	system.SaveFile(source, "content")
	system.ReplaceFile(source, destination)
	assert.String(t, "content", system.ReadFile(d, "destination"))
}
