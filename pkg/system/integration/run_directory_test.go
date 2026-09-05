package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/run"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunDirectory(t *testing.T) {
	windowsSkip(t)
	directory, e := filepath.EvalSymlinks(t.TempDir())
	errors.PanicOnError(e)
	r := run.New()
	r.Directory = directory
	output := r.Start("sh", "-c", "pwd")
	assert.String(t, directory, strings.TrimSpace(output))
}
