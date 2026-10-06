package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/snapshot"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestSnapshotReportsChangedAddedAndRemovedFiles(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(t, root, "go.mod", "module example\n")
	testutil.WriteFile(t, root, "a.go", "package example\n")
	testutil.WriteFile(t, root, "gone.go", "package example\n")
	testutil.WriteFile(t, root, "same.go", "package example\n")
	s := snapshot.Take(root)
	assert.Strings(t, nil, s.Changed(root))
	testutil.WriteFile(t, root, "a.go", "package other\n\n")
	testutil.WriteFile(t, root, "pkg/new.go", "package pkg\n")
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

func TestSnapshotIgnoresWhatTheModuleDoesNotBuild(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(t, root, "go.mod", "module example\n")
	s := snapshot.Take(root)
	testutil.WriteFile(t, root, "testdata/a.go", "package a\n")
	testutil.WriteFile(t, root, "_scratch/a.go", "package a\n")
	testutil.WriteFile(t, root, ".hidden/a.go", "package a\n")
	testutil.WriteFile(t, root, "vendor/a.go", "package a\n")
	testutil.WriteFile(t, root, "nested/go.mod", "module nested\n")
	testutil.WriteFile(t, root, "nested/a.go", "package a\n")
	testutil.WriteFile(t, root, constant.ReadmeFile, "readme\n")
	assert.Strings(t, nil, s.Changed(root))
	testutil.WriteFile(t, root, "go.sum", "sum\n")
	assert.Strings(t, []string{filepath.Join(root, "go.sum")}, s.Changed(root))
}
