package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/join"
	"testing"
)

func TestMoveCopyMovesContentAndRemovesSource(t *testing.T) {
	d := t.TempDir()
	source := join.Absolute(d, "source")
	system.SaveFile(source, "content")
	system.MoveCopy(source, join.Absolute(d, "destination"))
	assert.String(t, "content", system.ReadFile(d, "destination"))
	assert.False(t, system.FileExists(source))
}

func TestMoveCopyKeepsExecutableBit(t *testing.T) {
	d := t.TempDir()
	source := join.Absolute(d, "source")
	destination := join.Absolute(d, "destination")
	system.SaveFile(source, "content")
	system.Executable(source)
	system.MoveCopy(source, destination)
	assert.True(t, system.IsExecutable(destination))
}

func TestMoveCopyLogsNothing(t *testing.T) {
	d := t.TempDir()
	source := join.Absolute(d, "source")
	system.SaveFile(source, "content")
	assert.String(
		t,
		"",
		captureLog(
			func() { system.MoveCopy(source, join.Absolute(d, "destination")) },
		),
	)
}
