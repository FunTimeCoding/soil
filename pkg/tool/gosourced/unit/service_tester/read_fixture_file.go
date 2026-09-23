package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
	"path/filepath"
	"testing"
)

func ReadFixtureFile(
	t *testing.T,
	d string,
	path string,
) string {
	t.Helper()
	b, e := os.ReadFile(filepath.Join(d, path))
	assert.FatalOnError(t, e)

	return string(b)
}
