package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"testing"
)

func TestWalkSkippedTreeKnownNotRead(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(t, root, "doc/a.md", "# A\n")
	testutil.WriteFile(t, root, "tmp/note.txt", "draft\n")
	testutil.WriteFile(t, root, ".git/HEAD", "ref: refs/heads/main\n")
	testutil.WriteFile(t, root, "pkg/x/generated.go", "package x\n")
	repo, empty := lint.Walk(root, option.New("", false))
	assert.String(t, root, repo.Root)
	assert.String(t, "# A\n", repo.Files.ReadString("doc/a.md"))
	assert.True(t, repo.Files.Has("tmp/note.txt"))
	assert.String(t, "", repo.Files.ReadString("tmp/note.txt"))
	assert.Integer(t, 6, int(repo.Files.FileAt("tmp/note.txt").Size))
	assert.Boolean(t, false, repo.Files.Has(".git/HEAD"))
	assert.String(t, "package x\n", repo.Files.ReadString("pkg/x/generated.go"))
	assert.Strings(t, nil, empty)
}

func TestWalkStopsAtNestedModule(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(t, root, "go.mod", "module alfa\n")
	testutil.WriteFile(t, root, "doc/a.md", "# A\n")
	testutil.WriteFile(t, root, "graft/bravo/go.mod", "module bravo\n")
	testutil.WriteFile(t, root, "graft/bravo/doc/b.md", "# Bravo\n")
	repo, _ := lint.Walk(root, option.New("", false))
	assert.True(t, repo.Files.Has("doc/a.md"))
	assert.Boolean(t, false, repo.Files.Has("graft/bravo/go.mod"))
	assert.Boolean(t, false, repo.Files.Has("graft/bravo/doc/b.md"))
	nested, _ := lint.Walk(
		filepath.Join(root, "graft", "bravo"),
		option.New("", false),
	)
	assert.True(t, nested.Files.Has("doc/b.md"))
}

func TestWalkEmptyDirectoriesOutsideSkips(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(t, root, "doc/a.md", "# A\n")
	system.MakeDirectory(filepath.Join(root, "doc", "empty"))
	system.MakeDirectory(filepath.Join(root, "tmp", "empty"))
	system.MakeDirectory(filepath.Join(root, ".git", "empty"))
	_, empty := lint.Walk(root, option.New("", false))
	assert.Strings(t, []string{"doc/empty"}, empty)
}
