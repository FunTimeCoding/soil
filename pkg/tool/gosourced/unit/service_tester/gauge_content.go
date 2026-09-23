package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
	"path/filepath"
	"testing"
)

func GaugeContent(
	t *testing.T,
	directory string,
) string {
	t.Helper()
	b, e := os.ReadFile(filepath.Join(directory, "pkg/gauge/run.go"))
	assert.FatalOnError(t, e)

	return string(b)
}
