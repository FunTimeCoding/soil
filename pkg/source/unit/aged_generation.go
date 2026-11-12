package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func agedGeneration(
	t *testing.T,
	root string,
	name string,
	touched time.Time,
) {
	t.Helper()
	directory := filepath.Join(root, name)
	assert.FatalOnError(t, os.MkdirAll(directory, 0755))
	marker := filepath.Join(directory, constant.GenerationMarker)
	assert.FatalOnError(t, os.WriteFile(marker, nil, 0644))
	assert.FatalOnError(t, os.Chtimes(marker, touched, touched))
}
