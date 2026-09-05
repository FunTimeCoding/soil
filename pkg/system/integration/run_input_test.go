package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/run"
	"strings"
	"testing"
)

func TestRunInput(t *testing.T) {
	windowsSkip(t)
	r := run.New()
	r.Input = strings.NewReader("first\nsecond\n")
	output := r.Start("cat")
	assert.True(t, r.Error == nil)
	assert.String(t, "first\nsecond\n", output)
	assert.String(t, "first\nsecond\n", r.OutputString)
}
