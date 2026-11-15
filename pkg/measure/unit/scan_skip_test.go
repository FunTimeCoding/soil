package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/measure"
	"github.com/funtimecoding/soil/pkg/measure/option"
	"github.com/funtimecoding/soil/pkg/measure/registry"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
	"testing"
)

func TestScanSkipByNameEverywhere(t *testing.T) {
	root := t.TempDir()
	system_tester.Write(t, root, "tmp/a.go", "package a\n")
	system_tester.Write(t, root, "pkg/tmp/b.go", "package b\n")
	system_tester.Write(t, root, "pkg/keep.go", "package keep\n")
	o := option.New()
	o.Paths = []string{root}
	o.Skips = []string{"tmp"}
	r := measure.Scan(o, registry.NewDefault())
	assert.Count(t, 1, r.Files)
	assert.String(t, filepath.Join(root, "pkg/keep.go"), r.Files[0].Path)
}

func TestScanSkipByPathPrefixOnly(t *testing.T) {
	root := t.TempDir()
	system_tester.Write(t, root, "tmp/a.go", "package a\n")
	system_tester.Write(t, root, "pkg/tmp/b.go", "package b\n")
	system_tester.Write(t, root, "tmpfile.go", "package c\n")
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
	system_tester.Write(t, root, "pkg/tmp/b.go", "package b\n")
	system_tester.Write(t, root, "pkg/keep.go", "package keep\n")
	o := option.New()
	o.Paths = []string{root}
	o.Skips = []string{"pkg/tmp/"}
	r := measure.Scan(o, registry.NewDefault())
	assert.Count(t, 1, r.Files)
	assert.String(t, filepath.Join(root, "pkg/keep.go"), r.Files[0].Path)
}
