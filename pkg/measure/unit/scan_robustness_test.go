package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"path/filepath"
	"testing"
)

func TestScanHonorsIgnore(t *testing.T) {
	root := t.TempDir()
	write(root, ".gitignore", "build/\n*.log\n")
	write(root, "keep.go", "package keep\n")
	write(root, "build/out.go", "package out\n")
	write(root, "notes.log", "line\n")
	write(root, "nested/.gitignore", "secret.go\n")
	write(root, "nested/secret.go", "package secret\n")
	write(root, "nested/open.go", "package open\n")
	r := scan(root)
	assert.Count(t, 4, r.Files)
	assert.Count(t, 0, r.Unplaced)
	v := r.ByPath()
	assert.String(t, filepath.Join(root, ".gitignore"), v[0].Path)
	assert.String(t, filepath.Join(root, "keep.go"), v[1].Path)
	assert.String(t, filepath.Join(root, "nested/.gitignore"), v[2].Path)
	assert.String(t, filepath.Join(root, "nested/open.go"), v[3].Path)
}

func TestScanSkipsBinary(t *testing.T) {
	root := t.TempDir()
	write(root, "blob.go", "package main\x00\x01\x02")
	write(root, "text.go", "package main\n")
	r := scan(root)
	assert.Count(t, 1, r.Files)
	assert.Strings(t, []string{filepath.Join(root, "blob.go")}, r.Skipped)
}

func TestScanSkipsUnreadable(t *testing.T) {
	root := t.TempDir()
	write(root, "text.go", "package main\n")
	write(root, "locked.go", "package main\n")
	unreadable(t, filepath.Join(root, "locked.go"))
	r := scan(root)
	assert.Count(t, 1, r.Files)
	assert.Strings(t, []string{filepath.Join(root, "locked.go")}, r.Skipped)
}
