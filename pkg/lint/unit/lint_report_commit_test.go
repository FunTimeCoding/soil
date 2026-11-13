package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"testing"
)

func TestReportFixRefusedWhenFixedFileChanges(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(
		t,
		root,
		"pkg/alfa/alfa.go",
		"package alfa\n\nfunc Bravo() int {\n\tcharlie := 1\n\treturn charlie\n}\n",
	)
	r := reportFixing(
		t,
		root,
		func() {
			testutil.WriteFile(
				t,
				root,
				"pkg/alfa/alfa.go",
				"package alfa\n\nfunc Bravo() int {\n\treturn 2\n}\n",
			)
		},
	)
	assert.String(
		t,
		"package alfa\n\nfunc Bravo() int {\n\treturn 2\n}\n",
		testutil.ReadFile(t, filepath.Join(root, "pkg/alfa/alfa.go")),
	)

	for _, c := range r.Entries {
		assert.False(t, c.Fixed)
	}

	last := r.Entries[len(r.Entries)-1]
	assert.String(t, "pkg/alfa/alfa.go", last.Path)
	assert.String(
		t,
		"changed since it was read - nothing written, run again",
		last.Text,
	)
}

func TestReportFixCommitsPastChangeToOtherFile(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(
		t,
		root,
		"pkg/alfa/alfa.go",
		"package alfa\n\nfunc Bravo() int {\n\tcharlie := 1\n\treturn charlie\n}\n",
	)
	testutil.WriteFile(
		t,
		root,
		"pkg/delta/delta.go",
		"package delta\n\nfunc Echo() int {\n\treturn 1\n}\n",
	)
	reportFixing(
		t,
		root,
		func() {
			testutil.WriteFile(t, root, "pkg/delta/delta.go", "package delta\n")
		},
	)
	assert.String(
		t,
		"package alfa\n\nfunc Bravo() int {\n\tcharlie := 1\n\n\treturn charlie\n}\n",
		testutil.ReadFile(t, filepath.Join(root, "pkg/alfa/alfa.go")),
	)
}

func TestReportFixKeepsEmptyFileFilledBeforeCommit(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(t, root, "pkg/alfa/alfa.go", "package alfa\n")
	testutil.WriteFile(t, root, "pkg/alfa/empty.go", "")
	reportFixing(
		t,
		root,
		func() {
			testutil.WriteFile(t, root, "pkg/alfa/empty.go", "package alfa\n")
		},
	)
	assert.True(t, system.FileExists(filepath.Join(root, "pkg/alfa/empty.go")))
}

func TestReportFixRemovesEmptyFile(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(t, root, "pkg/alfa/alfa.go", "package alfa\n")
	testutil.WriteFile(t, root, "pkg/alfa/empty.go", "")
	reportFixing(t, root, nil)
	assert.False(t, system.FileExists(filepath.Join(root, "pkg/alfa/empty.go")))
}
