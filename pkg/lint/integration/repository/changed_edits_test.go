package repository

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git/repository_tester"
	"github.com/funtimecoding/soil/pkg/lint/reach"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
	"testing"
)

func TestOnlyEditsGitStillSeesChangedAreInReach(t *testing.T) {
	root, e := filepath.EvalSymlinks(t.TempDir())
	assert.FatalOnError(t, e)
	system_tester.New(t).Write(root, "alfa.go", "package alfa\n")
	system_tester.New(t).Write(root, "bravo.go", "package alfa\n")
	system_tester.New(t).Write(root, "charlie.go", "package alfa\n")
	repository_tester.Identified(root, "init", "-q")
	repository_tester.Identified(root, "add", "--all")
	repository_tester.Identified(root, "commit", "--quiet", "--message=alfa")
	system_tester.New(t).Write(
		root,
		"bravo.go",
		"package alfa\n\nfunc Bravo() {}\n",
	)
	system_tester.New(t).Write(root, "delta.go", "package alfa\n")
	repository_tester.Identified(root, "mv", "charlie.go", "echo.go")
	changed, f := reach.ChangedEdits(
		root,
		reach.NewEdits(
			[]string{
				filepath.Join(root, "alfa.go"),
				filepath.Join(root, "bravo.go"),
				filepath.Join(root, "charlie.go"),
				filepath.Join(root, "delta.go"),
				filepath.Join(root, "echo.go"),
			},
		),
	)
	assert.FatalOnError(t, f)
	assert.Strings(
		t,
		[]string{
			filepath.Join(root, "bravo.go"),
			filepath.Join(root, "charlie.go"),
			filepath.Join(root, "delta.go"),
			filepath.Join(root, "echo.go"),
		},
		changed,
	)
}

func TestGlobReachesChangedFilesItMatches(t *testing.T) {
	root, e := filepath.EvalSymlinks(t.TempDir())
	assert.FatalOnError(t, e)
	system_tester.New(t).Write(root, "alfa.md", "alfa\n")
	system_tester.New(t).Write(root, "bravo.md", "bravo\n")
	system_tester.New(t).Write(root, "charlie.go", "package charlie\n")
	repository_tester.Identified(root, "init", "-q")
	repository_tester.Identified(root, "add", "--all")
	repository_tester.Identified(root, "commit", "--quiet", "--message=alfa")
	system_tester.New(t).Write(root, "bravo.md", "bravo changed\n")
	system_tester.New(t).Write(root, "delta.md", "delta\n")
	system_tester.New(t).Write(
		root,
		"charlie.go",
		"package charlie\n\nfunc C() {}\n",
	)
	edits := reach.NewEdits(nil)
	edits.Patterns = []string{filepath.Join(root, "*.md")}
	changed, f := reach.ChangedEdits(root, edits)
	assert.FatalOnError(t, f)
	assert.Strings(
		t,
		[]string{
			filepath.Join(root, "bravo.md"),
			filepath.Join(root, "delta.md"),
		},
		changed,
	)
}
