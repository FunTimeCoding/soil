package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/measure"
	"github.com/funtimecoding/soil/pkg/measure/option"
	"github.com/funtimecoding/soil/pkg/measure/registry"
	"path/filepath"
	"testing"
)

func TestScanSkipByNameEverywhere(t *testing.T) {
	root := t.TempDir()
	write(root, "tmp/a.go", "package a\n")
	write(root, "pkg/tmp/b.go", "package b\n")
	write(root, "pkg/keep.go", "package keep\n")
	o := option.New()
	o.Paths = []string{root}
	o.Skips = []string{"tmp"}
	r := measure.Scan(o, registry.NewDefault())
	assert.Count(t, 1, r.Files)
	assert.String(t, filepath.Join(root, "pkg/keep.go"), r.Files[0].Path)
}

func TestScanSkipByPathPrefixOnly(t *testing.T) {
	root := t.TempDir()
	write(root, "tmp/a.go", "package a\n")
	write(root, "pkg/tmp/b.go", "package b\n")
	write(root, "tmpfile.go", "package c\n")
	o := option.New()
	o.Paths = []string{root}
	o.Skips = []string{"tmp/"}
	r := measure.Scan(o, registry.NewDefault())
	v := r.ByPath()
	assert.Count(t, 2, v)
	assert.String(t, filepath.Join(root, "pkg/tmp/b.go"), v[0].Path)
	assert.String(t, filepath.Join(root, "tmpfile.go"), v[1].Path)
}

func TestScanSkipNestedPathPrefix(t *testing.T) {
	root := t.TempDir()
	write(root, "pkg/tmp/b.go", "package b\n")
	write(root, "pkg/keep.go", "package keep\n")
	o := option.New()
	o.Paths = []string{root}
	o.Skips = []string{"pkg/tmp/"}
	r := measure.Scan(o, registry.NewDefault())
	assert.Count(t, 1, r.Files)
	assert.String(t, filepath.Join(root, "pkg/keep.go"), r.Files[0].Path)
}
