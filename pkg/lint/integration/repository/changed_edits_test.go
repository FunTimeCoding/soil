package repository

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/reach"
	"path/filepath"
	"testing"
)

func TestOnlyEditsGitStillSeesChangedAreInReach(t *testing.T) {
	root, e := filepath.EvalSymlinks(t.TempDir())
	assert.FatalOnError(t, e)
	testutil.WriteFile(t, root, "alfa.go", "package alfa\n")
	testutil.WriteFile(t, root, "bravo.go", "package alfa\n")
	testutil.WriteFile(t, root, "charlie.go", "package alfa\n")
	git(t, root, "init", "-q")
	git(t, root, "add", "--all")
	git(t, root, "commit", "--quiet", "--message=alfa")
	testutil.WriteFile(t, root, "bravo.go", "package alfa\n\nfunc Bravo() {}\n")
	testutil.WriteFile(t, root, "delta.go", "package alfa\n")
	git(t, root, "mv", "charlie.go", "echo.go")
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
	testutil.WriteFile(t, root, "alfa.md", "alfa\n")
	testutil.WriteFile(t, root, "bravo.md", "bravo\n")
	testutil.WriteFile(t, root, "charlie.go", "package charlie\n")
	git(t, root, "init", "-q")
	git(t, root, "add", "--all")
	git(t, root, "commit", "--quiet", "--message=alfa")
	testutil.WriteFile(t, root, "bravo.md", "bravo changed\n")
	testutil.WriteFile(t, root, "delta.md", "delta\n")
	testutil.WriteFile(
		t,
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
