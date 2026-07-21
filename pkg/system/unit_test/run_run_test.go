package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"runtime"
	"testing"
)

func TestRunStart(t *testing.T) {
	r1 := run.New()
	assert.True(t, r1.Panic)
	r1.Panic = false
	r1.Start("echo", "test")
	assert.True(t, r1.Error == nil)
	assert.String(t, "test\n", r1.OutputString)
	assert.String(t, "", r1.ErrorString)
	r2 := run.New()
	r2.Panic = false
	r2.Start("nonexistent")
	assert.True(t, r2.Error != nil)
	assert.String(t, "", r2.OutputString)
	assert.String(t, "", r2.ErrorString)

	switch runtime.GOOS {
	case constant.Windows:
		assert.String(
			t,
			`exec: "nonexistent": executable file not found in %PATH%`,
			r2.Error.Error(),
		)
	default:
		assert.String(
			t,
			`exec: "nonexistent": executable file not found in $PATH`,
			r2.Error.Error(),
		)
	}
}
