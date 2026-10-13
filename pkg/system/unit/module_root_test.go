package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"testing"
)

func TestModuleRoot(t *testing.T) {
	outer := t.TempDir()
	testutil.WriteFile(t, outer, "go.mod", "module outer\n")
	root := filepath.Join(outer, "repository")
	testutil.WriteFile(t, root, "go.mod", "module alfa\n")
	testutil.WriteFile(t, root, "pkg/alfa/alfa.go", "package alfa\n")
	testutil.WriteFile(t, root, "graft/bravo/go.mod", "module bravo\n")
	testutil.WriteFile(
		t,
		root,
		"graft/bravo/pkg/bravo/bravo.go",
		"package bravo\n",
	)
	testutil.WriteFile(t, root, "doc/a.md", "# A\n")
	nested := filepath.Join(root, "graft", "bravo")
	assert.String(t, root, system.ModuleRoot(root, root))
	assert.String(
		t,
		root,
		system.ModuleRoot(filepath.Join(root, "pkg", "alfa"), root),
	)
	assert.String(t, nested, system.ModuleRoot(nested, root))
	assert.String(
		t,
		nested,
		system.ModuleRoot(filepath.Join(nested, "pkg", "bravo"), root),
	)
	assert.String(
		t,
		root,
		system.ModuleRoot(filepath.Join(root, "graft"), root),
	)
}

func TestModuleRootWithoutModule(t *testing.T) {
	outer := t.TempDir()
	testutil.WriteFile(t, outer, "go.mod", "module outer\n")
	root := filepath.Join(outer, "repository")
	testutil.WriteFile(t, root, "notes/a.md", "# A\n")
	assert.String(
		t,
		root,
		system.ModuleRoot(filepath.Join(root, "notes"), root),
	)
}
