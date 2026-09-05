package integration

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/run"
	"strings"
	"testing"
)

func TestRunInputWithWriters(t *testing.T) {
	windowsSkip(t)
	var stderr bytes.Buffer
	r := run.New()
	r.NoPanic()
	r.Input = strings.NewReader("banana\napple\ncherry\n")
	r.Writers(nil, &stderr)
	output := r.Start("sort")
	assert.True(t, r.Error == nil)
	assert.String(t, "apple\nbanana\ncherry\n", output)
	assert.String(t, "", stderr.String())
}
