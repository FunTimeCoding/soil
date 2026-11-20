package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/source/warm_tree"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestWarmTreeHoldsEntriesButGit(t *testing.T) {
	root := t.TempDir()
	system_tester.New(t).Write(root, "alfa.txt", "bravo\n")
	system_tester.New(t).Write(root, "charlie/delta.go", "package charlie\n")
	system_tester.New(t).Write(root, ".git/config", "echo\n")
	system.MakeDirectory(filepath.Join(root, "empty"))
	assert.FatalOnError(
		t,
		os.Symlink(
			filepath.Join(root, "charlie"),
			filepath.Join(root, "linked"),
		),
	)
	assert.FatalOnError(
		t,
		os.Symlink(
			filepath.Join(root, "alfa.txt"),
			filepath.Join(root, "file"),
		),
	)
	tree := warm_tree.New(root)
	assert.Strings(t, []string{"alfa.txt", "charlie/delta.go"}, tree.Files())
	assert.Strings(
		t,
		[]string{"alfa.txt", "charlie", "empty", "linked"},
		tree.Children(soil.CurrentDirectory),
	)
	empty, _ := tree.Stat("empty")
	assert.True(t, empty.Directory)
	linked, _ := tree.Stat("linked")
	assert.True(t, linked.Directory)
	_, found := tree.Stat(".git/config")
	assert.False(t, found)
}

func TestWarmTreeReadsOnceThenFromMemory(t *testing.T) {
	root := t.TempDir()
	system_tester.New(t).Write(root, "alfa.txt", "bravo\n")
	tree := warm_tree.New(root)
	first, e := tree.Read("alfa.txt")
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, os.Remove(filepath.Join(root, "alfa.txt")))
	second, f := tree.Read("./alfa.txt")
	assert.FatalOnError(t, f)
	assert.String(t, "bravo\n", string(first))
	assert.String(t, "bravo\n", string(second))
}

func TestWarmTreeRefreshNamesWhatChanged(t *testing.T) {
	root := t.TempDir()
	system_tester.New(t).Write(root, "alfa.txt", "bravo\n")
	system_tester.New(t).Write(root, "charlie/delta.go", "package charlie\n")
	system_tester.New(t).Write(root, "echo/foxtrot.go", "package echo\n")
	tree := warm_tree.New(root)
	_, e := tree.Read("alfa.txt")
	assert.FatalOnError(t, e)
	system_tester.New(t).Write(root, "alfa.txt", "golf hotel\n")
	system_tester.New(t).Write(root, "charlie/india.go", "package charlie\n")
	assert.FatalOnError(t, os.Remove(filepath.Join(root, "echo/foxtrot.go")))
	changed := tree.Refresh()
	assert.Strings(
		t,
		[]string{
			"alfa.txt",
			"charlie",
			"charlie/india.go",
			"echo",
			"echo/foxtrot.go",
		},
		changed,
	)
	b, f := tree.Read("alfa.txt")
	assert.FatalOnError(t, f)
	assert.String(t, "golf hotel\n", string(b))
	assert.Strings(t, nil, tree.Refresh())
}

func TestWarmTreeIgnoreMatcherAgreesWithGit(t *testing.T) {
	root := t.TempDir()
	system_tester.New(t).Write(root, ".gitignore", "build/\n*.log\n!keep.log\n")
	system_tester.New(t).Write(root, ".git/info/exclude", "local/\n")
	system_tester.New(t).Write(root, "build/.gitignore", "!inside.txt\n")
	system_tester.New(t).Write(root, "build/inside.txt", "alfa\n")
	system_tester.New(t).Write(root, "local/.gitignore", "!secret.txt\n")
	system_tester.New(t).Write(root, "local/secret.txt", "bravo\n")
	system_tester.New(t).Write(root, "doc/.gitignore", "draft.md\n")
	system_tester.New(t).Write(root, "doc/draft.md", "charlie\n")
	system_tester.New(t).Write(root, "doc/final.md", "delta\n")
	system_tester.New(t).Write(root, "alfa.log", "echo\n")
	system_tester.New(t).Write(root, "keep.log", "foxtrot\n")
	paths := []string{
		"alfa.log",
		"build/inside.txt",
		"doc/draft.md",
		"doc/final.md",
		"keep.log",
		"local/secret.txt",
	}
	disk := git.IgnoreMatcher(root)
	warm := warm_tree.New(root).IgnoreMatcher()
	var expected []string
	var actual []string

	for _, p := range paths {
		if disk(p) {
			expected = append(expected, p)
		}

		if warm(p) {
			actual = append(actual, p)
		}
	}

	assert.Strings(
		t,
		[]string{
			"alfa.log",
			"build/inside.txt",
			"doc/draft.md",
			"local/secret.txt",
		},
		actual,
	)
	assert.Strings(t, expected, actual)
}
