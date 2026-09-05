package integration

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/run"
	"testing"
)

func TestRunWritersStdout(t *testing.T) {
	windowsSkip(t)
	var stdout bytes.Buffer
	r := run.New()
	r.Writers(&stdout, nil)
	output := r.Start("sh", "-c", "echo out; echo err >&2")
	assert.String(t, "out\n", stdout.String())
	assert.String(t, "", output)
	assert.String(t, "", r.OutputString)
	assert.String(t, "err\n", r.ErrorString)
}
