package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
	"testing"
)

func TestRepositoryExistsIndependentOfWorkingDirectory(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(t, root, "doc/a.md", "# A\n")
	testutil.WriteFile(t, root, "tmp/note.txt", "draft\n")
	system.MakeDirectory(filepath.Join(root, "empty"))
	system.MakeDirectory(filepath.Join(root, soil.ParentDirectory, "beside"))
	repo, _ := lint.Walk(root, option.New("", false))
	assert.String(t, filepath.Join(root, "doc"), repo.Absolute("doc"))
	assert.True(t, repo.Exists("doc"))
	assert.True(t, repo.Exists("doc/a.md"))
	assert.True(t, repo.Exists("tmp/note.txt"))
	assert.True(t, repo.Exists("empty"))
	assert.Boolean(t, false, repo.Exists("ghost"))
	assert.Boolean(t, false, repo.Exists(""))
	assert.True(t, repo.SiblingExists("../beside"))
	assert.Boolean(t, false, repo.SiblingExists("../ghost"))
	assert.Boolean(t, false, repo.SiblingExists(""))
	assert.Strings(t, nil, repo.Siblings)
	assert.Strings(t, nil, repo.ImplicitBases)
}

func TestRepositoryExistsAnswersFromWalk(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(t, root, "doc/a.md", "# A\n")
	system.MakeDirectory(filepath.Join(root, "real"))
	assert.FatalOnError(
		t,
		os.Symlink(filepath.Join(root, "real"), filepath.Join(root, "linked")),
	)
	repo, _ := lint.Walk(root, option.New("", false))
	system.MakeDirectory(filepath.Join(root, "late"))
	assert.Boolean(t, false, repo.Exists("late"))
	assert.True(t, repo.Exists("doc/"))
	assert.True(t, repo.Exists("doc/./"))
	assert.True(t, repo.Exists("linked"))
}

func TestRepositoryExistsIntoSiblingFromOneWalk(t *testing.T) {
	repo, sibling := siblingRepository(t)
	assert.True(t, repo.Exists("../bravo/pkg/charlie"))
	assert.True(t, repo.Exists("../bravo/pkg/empty"))
	assert.Boolean(t, false, repo.Exists("../bravo/ghost"))
	system.MakeDirectory(filepath.Join(sibling, "late"))
	assert.Boolean(t, false, repo.Exists("../bravo/late"))
}

func TestRepositoryLiteralInSiblingFromOneWalk(t *testing.T) {
	repo, _ := siblingRepository(t)
	assert.True(t, repo.Literal("../bravo/pkg", "delta echo"))
	assert.Boolean(t, false, repo.Literal("../bravo/pkg", "foxtrot golf"))
	assert.True(
		t,
		repo.Literal("../bravo/pkg/charlie/charlie.go", "delta echo"),
	)
}

func TestRepositoryExistsIntoNestedModule(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(t, root, "go.mod", "module alfa\n")
	testutil.WriteFile(t, root, "graft/bravo/go.mod", "module bravo\n")
	testutil.WriteFile(t, root, "graft/bravo/taskfile.yaml", "version: 3\n")
	testutil.WriteFile(t, root, "graft/charlie.txt", "charlie\n")
	repo, _ := lint.Walk(root, option.New("", false))
	assert.Strings(t, []string{"graft/bravo"}, repo.NestedModules)
	assert.Boolean(t, false, repo.Files.Has("graft/bravo/taskfile.yaml"))
	assert.True(t, repo.Exists("graft/bravo"))
	assert.True(t, repo.Exists("graft/bravo/go.mod"))
	assert.True(t, repo.Exists("graft/bravo/taskfile.yaml"))
	assert.Boolean(t, false, repo.Exists("graft/bravo/ghost.yaml"))
	assert.Boolean(t, false, repo.Exists("graft/bravoghost.txt"))
	assert.True(t, repo.Exists("graft/charlie.txt"))
}
