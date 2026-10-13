package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/git/unit/repository_tester"
	"path/filepath"
	"testing"
)

func TestModifiedFilesAtRoot(t *testing.T) {
	r := modifiedRepository(t)
	assert.Strings(
		t,
		[]string{"go.mod", "graft/bravo/go.mod", "graft/bravo/staged.md"},
		git.MustModifiedFiles(r.Clone),
	)
}

func TestModifiedFilesInSubdirectory(t *testing.T) {
	r := modifiedRepository(t)
	assert.Strings(
		t,
		[]string{"go.mod", "staged.md"},
		git.MustModifiedFiles(filepath.Join(r.Clone, "graft", "bravo")),
	)
}

func TestModifiedFilesStatusStaysRootRelative(t *testing.T) {
	r := modifiedRepository(t)
	s := git.Status(filepath.Join(r.Clone, "graft", "bravo"))
	_, okay := s["graft/bravo/go.mod"]
	assert.True(t, okay)
}

func TestModifiedFilesOutsideRepository(t *testing.T) {
	_, e := git.ModifiedFiles(t.TempDir())
	assert.Error(t, e)
}

func TestModifiedFilesInLinkedWorktree(t *testing.T) {
	r := repository_tester.New(t)
	r.Write("seed.md", "# Seed\n")
	r.Commit("seed")
	linked := filepath.Join(filepath.Dir(r.Clone), "linked")
	repository_tester.Command(r.Clone, "worktree", "add", linked)
	l := repository_tester.NewAt(linked)
	l.Write("seed.md", "# Changed\n")
	assert.Strings(t, []string{"seed.md"}, git.MustModifiedFiles(linked))
	assert.Strings(t, nil, git.MustModifiedFiles(r.Clone))
}
