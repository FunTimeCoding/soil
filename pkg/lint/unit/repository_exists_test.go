package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"testing"
)

func TestRepositoryExistsIndependentOfWorkingDirectory(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(t, root, "doc/a.md", "# A\n")
	testutil.WriteFile(t, root, "tmp/note.txt", "draft\n")
	system.MakeDirectory(filepath.Join(root, "empty"))
	system.MakeDirectory(
		filepath.Join(root, constant.ParentDirectory, "beside"),
	)
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
