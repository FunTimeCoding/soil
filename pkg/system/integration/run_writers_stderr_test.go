package integration

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/run"
	"testing"
)

func TestRunWritersStderr(t *testing.T) {
	windowsSkip(t)
	var stderr bytes.Buffer
	r := run.New()
	r.Writers(nil, &stderr)
	output := r.Start("sh", "-c", "echo out; echo err >&2")
	assert.String(t, "out\n", output)
	assert.String(t, "out\n", r.OutputString)
	assert.String(t, "err\n", stderr.String())
	assert.String(t, "", r.ErrorString)
}
