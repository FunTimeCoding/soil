package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/source/snapshot"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestSnapshotReportsChangedAddedAndRemovedFiles(t *testing.T) {
	root := t.TempDir()
	system_tester.New(t).Write(root, "go.mod", "module example\n")
	system_tester.New(t).Write(root, "a.go", "package example\n")
	system_tester.New(t).Write(root, "gone.go", "package example\n")
	system_tester.New(t).Write(root, "same.go", "package example\n")
	s := snapshot.Take(root)
	assert.Strings(t, nil, s.Changed(root))
	system_tester.New(t).Write(root, "a.go", "package other\n\n")
	system_tester.New(t).Write(root, "pkg/new.go", "package pkg\n")
	assert.FatalOnError(t, os.Remove(filepath.Join(root, "gone.go")))
	later := time.Now().Add(time.Hour)
	assert.FatalOnError(
		t,
		os.Chtimes(filepath.Join(root, "same.go"), later, later),
	)
	assert.Strings(
		t,
		[]string{
			filepath.Join(root, "a.go"),
			filepath.Join(root, "gone.go"),
			filepath.Join(root, "pkg/new.go"),
			filepath.Join(root, "same.go"),
		},
		s.Changed(root),
	)
}

func TestSnapshotIgnoresWhatModuleDoesNotBuild(t *testing.T) {
	root := t.TempDir()
	system_tester.New(t).Write(root, "go.mod", "module example\n")
	s := snapshot.Take(root)
	system_tester.New(t).Write(root, "testdata/a.go", "package a\n")
	system_tester.New(t).Write(root, "_scratch/a.go", "package a\n")
	system_tester.New(t).Write(root, ".hidden/a.go", "package a\n")
	system_tester.New(t).Write(root, "vendor/a.go", "package a\n")
	system_tester.New(t).Write(root, "nested/go.mod", "module nested\n")
	system_tester.New(t).Write(root, "nested/a.go", "package a\n")
	system_tester.New(t).Write(root, soil.ReadmeFile, "readme\n")
	assert.Strings(t, nil, s.Changed(root))
	system_tester.New(t).Write(root, "go.sum", "sum\n")
	assert.Strings(t, []string{filepath.Join(root, "go.sum")}, s.Changed(root))
}
