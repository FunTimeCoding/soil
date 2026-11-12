package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestClientPrunesStaleIndexGeneration(t *testing.T) {
	directory := t.TempDir()
	assert.FatalOnError(
		t,
		os.MkdirAll(filepath.Join(directory, "go1.0-0123456789ab"), 0755),
	)
	removed, e := mountedIndex(t, directory).PostIndexPrune(t.Context())
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"go1.0-0123456789ab"}, removed)
	_, f := os.Stat(filepath.Join(directory, "go1.0-0123456789ab"))
	assert.True(t, os.IsNotExist(f))
}
