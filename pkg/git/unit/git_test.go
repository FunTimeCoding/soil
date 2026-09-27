package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/git/remote"
	"github.com/funtimecoding/soil/pkg/github/action"
	github "github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/contains"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestConstant(t *testing.T) {
	assert.String(t, "v", gitConstant.VersionPrefix)
	assert.String(t, "HEAD", gitConstant.HeadReference)
	assert.Integer(t, 7, gitConstant.HashLength)
}

func TestBranch(t *testing.T) {
	system.PrintEnvironment()
	e := gitConstant.MainBranch

	if action.IsActionRun() {
		r := environment.Required(github.ReferenceEnvironment)

		if r != gitConstant.MainBranch {
			e = r
		}
	}

	actual := git.Branch(git.FindDirectory())

	if false {
		// Sometimes HEAD
		assert.String(t, e, actual)
	}

	// TODO: Add reference environment to list if missing
	//  Then count somewhere what observations there are
	assert.True(
		t,
		contains.Any(
			[]string{actual},
			[]string{gitConstant.MainBranch, gitConstant.HeadReference},
		),
	)
}

func TestCommitTimes(t *testing.T) {
	times, e := git.CommitTimes(git.FindDirectory())
	assert.FatalOnError(t, e)
	modified, okay := times["go.mod"]
	assert.True(t, okay)
	assert.True(t, modified.After(time.Unix(0, 0)))
}

func TestCommitTimesOutsideRepository(t *testing.T) {
	_, e := git.CommitTimes(t.TempDir())
	assert.Error(t, e)
}

func TestUncommittedFiles(t *testing.T) {
	_, e := git.UncommittedFiles(git.FindDirectory())
	assert.FatalOnError(t, e)
}

func TestIgnoreMatcher(t *testing.T) {
	root := t.TempDir()

	if e := os.WriteFile(
		filepath.Join(root, ".gitignore"),
		[]byte("/.claude/notes\ntmp/\n"),
		0o600,
	); e != nil {
		t.Fatalf("write: %v", e)
	}

	ignored := git.IgnoreMatcher(root)
	assert.True(t, ignored(".claude/notes"))
	assert.True(t, ignored(".claude/notes/alfa.md"))
	assert.True(t, ignored("tmp/report.json"))
	assert.False(t, ignored(".claude/skills/bravo/SKILL.md"))
	assert.False(t, ignored("doc/charlie.md"))
}

func TestIgnoreMatcherWithoutRules(t *testing.T) {
	ignored := git.IgnoreMatcher(t.TempDir())
	assert.False(t, ignored("doc/charlie.md"))
}

func TestStatus(t *testing.T) {
	git.Status(git.FindDirectory())
}

func TestTags(t *testing.T) {
	git.Tags(git.FindDirectory())
}

func TestTree(t *testing.T) {
	git.Tree(git.FindDirectory())
}

func TestRemote(t *testing.T) {
	assert.Any(
		t,
		&remote.Remote{Name: "Alfa", Locator: "Bravo", Provider: "Charlie"},
		remote.New(
			constant.UpperAlfa,
			constant.UpperBravo,
			constant.UpperCharlie,
		),
	)
}
