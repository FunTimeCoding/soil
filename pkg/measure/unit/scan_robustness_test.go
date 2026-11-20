package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/measure/unit/measure_tester"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
	"testing"
)

func TestScanHonorsIgnore(t *testing.T) {
	root := t.TempDir()
	system_tester.New(t).Write(root, ".gitignore", "build/\n*.log\n")
	system_tester.New(t).Write(root, "keep.go", "package keep\n")
	system_tester.New(t).Write(root, "build/out.go", "package out\n")
	system_tester.New(t).Write(root, "notes.log", "line\n")
	system_tester.New(t).Write(root, "nested/.gitignore", "secret.go\n")
	system_tester.New(t).Write(root, "nested/secret.go", "package secret\n")
	system_tester.New(t).Write(root, "nested/open.go", "package open\n")
	r := measure_tester.Scan(root)
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
	system_tester.New(t).Write(root, "blob.go", "package main\x00\x01\x02")
	system_tester.New(t).Write(root, "text.go", "package main\n")
	r := measure_tester.Scan(root)
	assert.Count(t, 1, r.Files)
	assert.Strings(t, []string{filepath.Join(root, "blob.go")}, r.Skipped)
}

func TestScanSkipsUnreadable(t *testing.T) {
	root := t.TempDir()
	system_tester.New(t).Write(root, "text.go", "package main\n")
	system_tester.New(t).Write(root, "locked.go", "package main\n")
	measure_tester.Unreadable(t, filepath.Join(root, "locked.go"))
	r := measure_tester.Scan(root)
	assert.Count(t, 1, r.Files)
	assert.Strings(t, []string{filepath.Join(root, "locked.go")}, r.Skipped)
}
