package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/unit/lint_tester"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
	"testing"
)

func TestReportFixRefusedWhenFixedFileChanges(t *testing.T) {
	root := t.TempDir()
	system_tester.New(t).Write(
		root,
		"pkg/alfa/alfa.go",
		"package alfa\n\nfunc Bravo() int {\n\tcharlie := 1\n\treturn charlie\n}\n",
	)
	r := lint_tester.New(t).FixingReport(
		root,
		func() {
			system_tester.New(t).Write(
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
	system_tester.New(t).Write(
		root,
		"pkg/alfa/alfa.go",
		"package alfa\n\nfunc Bravo() int {\n\tcharlie := 1\n\treturn charlie\n}\n",
	)
	system_tester.New(t).Write(
		root,
		"pkg/delta/delta.go",
		"package delta\n\nfunc Echo() int {\n\treturn 1\n}\n",
	)
	lint_tester.New(t).FixingReport(
		root,
		func() {
			system_tester.New(t).Write(
				root,
				"pkg/delta/delta.go",
				"package delta\n",
			)
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
	system_tester.New(t).Write(root, "pkg/alfa/alfa.go", "package alfa\n")
	system_tester.New(t).Write(root, "pkg/alfa/empty.go", "")
	lint_tester.New(t).FixingReport(
		root,
		func() {
			system_tester.New(t).Write(
				root,
				"pkg/alfa/empty.go",
				"package alfa\n",
			)
		},
	)
	assert.True(t, system.FileExists(filepath.Join(root, "pkg/alfa/empty.go")))
}

func TestReportFixRemovesEmptyFile(t *testing.T) {
	root := t.TempDir()
	system_tester.New(t).Write(root, "pkg/alfa/alfa.go", "package alfa\n")
	system_tester.New(t).Write(root, "pkg/alfa/empty.go", "")
	lint_tester.New(t).FixingReport(root, nil)
	assert.False(t, system.FileExists(filepath.Join(root, "pkg/alfa/empty.go")))
}
