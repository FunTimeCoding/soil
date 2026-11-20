package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
	"testing"
)

func TestModuleRoot(t *testing.T) {
	outer := t.TempDir()
	system_tester.New(t).Write(outer, "go.mod", "module outer\n")
	root := filepath.Join(outer, "repository")
	system_tester.New(t).Write(root, "go.mod", "module alfa\n")
	system_tester.New(t).Write(root, "pkg/alfa/alfa.go", "package alfa\n")
	system_tester.New(t).Write(root, "graft/bravo/go.mod", "module bravo\n")
	system_tester.New(t).Write(
		root,
		"graft/bravo/pkg/bravo/bravo.go",
		"package bravo\n",
	)
	system_tester.New(t).Write(root, "doc/a.md", "# A\n")
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
	system_tester.New(t).Write(outer, "go.mod", "module outer\n")
	root := filepath.Join(outer, "repository")
	system_tester.New(t).Write(root, "notes/a.md", "# A\n")
	assert.String(
		t,
		root,
		system.ModuleRoot(filepath.Join(root, "notes"), root),
	)
}
