package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"testing"
)

func TestFirstFile(t *testing.T) {
	root := t.TempDir()
	second := filepath.Join(root, "second.yaml")
	errors.PanicClose(system.Create(second))
	assert.String(t, "", system.FirstFile())
	assert.String(t, "", system.FirstFile(filepath.Join(root, "missing.yaml")))
	assert.String(
		t,
		second,
		system.FirstFile(filepath.Join(root, "first.yaml"), second),
	)
	assert.String(t, "", system.FirstFile(root))
}
