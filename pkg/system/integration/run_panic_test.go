package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"testing"
)

func TestRunPanicMode(t *testing.T) {
	windowsSkip(t)
	var recovered any
	func() {
		defer func() { recovered = recover() }()
		run.New().Start(constant.Shell, constant.ShellCommand, "exit 7")
	}()
	assert.NotNil(t, recovered)
}
