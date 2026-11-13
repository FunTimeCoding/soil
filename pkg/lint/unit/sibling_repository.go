package unit

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/repository"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"testing"
)

func siblingRepository(t *testing.T) (*repository.Repository, string) {
	t.Helper()
	parent := t.TempDir()
	root := filepath.Join(parent, "alfa")
	sibling := filepath.Join(parent, "bravo")
	testutil.WriteFile(
		t,
		root,
		"go.mod",
		"module example.org/alfa\n\nrequire example.org/bravo v1.0.0\n",
	)
	testutil.WriteFile(t, root, "doc/a.md", "# A\n")
	testutil.WriteFile(
		t,
		sibling,
		"pkg/charlie/charlie.go",
		"package charlie\n\nconst Charlie = \"delta echo\"\n",
	)
	system.MakeDirectory(filepath.Join(sibling, "pkg/empty"))
	repo, _ := lint.Walk(root, option.New("", false))

	return repo, sibling
}
