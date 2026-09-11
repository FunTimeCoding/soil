package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
	"testing"
)

func TestFindDirectoryUp(t *testing.T) {
	root := t.TempDir()
	assert.FatalOnError(t, os.MkdirAll(filepath.Join(root, ".git"), 0755))
	nested := filepath.Join(root, "alpha", "bravo")
	assert.FatalOnError(t, os.MkdirAll(nested, 0755))
	assert.String(t, root, system.FindDirectoryUp(nested, ".git"))
	assert.String(t, "", system.FindDirectoryUp(nested, "absent.marker"))
}

func TestFindDirectoryUpRelative(t *testing.T) {
	root := t.TempDir()
	nested := filepath.Join(root, "alpha", "bravo")
	assert.FatalOnError(t, os.MkdirAll(nested, 0755))
	assert.FatalOnError(
		t,
		os.MkdirAll(filepath.Join(root, "alpha", ".git"), 0755),
	)
	previous := system.WorkDirectory()
	assert.FatalOnError(t, os.Chdir(root))

	defer func() { assert.FatalOnError(t, os.Chdir(previous)) }()
	assert.String(
		t,
		filepath.Join("alpha"),
		system.FindDirectoryUp(filepath.Join("alpha", "bravo"), ".git"),
	)
}
