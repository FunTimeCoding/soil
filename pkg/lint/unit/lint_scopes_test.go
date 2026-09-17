package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"path/filepath"
	"testing"
)

func TestScopesNone(t *testing.T) {
	root := t.TempDir()
	scopes, e := lint.Scopes(root, root, nil)
	assert.Nil(t, e)
	assert.Strings(t, nil, scopes)
}

func TestScopesFromRoot(t *testing.T) {
	root := scopeTree(t)
	scopes, e := lint.Scopes(root, root, []string{"doc/ai/spec/naming.md"})
	assert.Nil(t, e)
	assert.Strings(t, []string{"doc/ai/spec/naming.md"}, scopes)
}

func TestScopesFromSubdirectory(t *testing.T) {
	root := scopeTree(t)
	scopes, e := lint.Scopes(
		root,
		filepath.Join(root, "doc", "ai"),
		[]string{"spec", "../../pkg"},
	)
	assert.Nil(t, e)
	assert.Strings(t, []string{"doc/ai/spec", "pkg"}, scopes)
}

func TestScopesAbsolute(t *testing.T) {
	root := scopeTree(t)
	scopes, e := lint.Scopes(
		root,
		t.TempDir(),
		[]string{filepath.Join(root, constant.PackageDirectory)},
	)
	assert.Nil(t, e)
	assert.Strings(t, []string{"pkg"}, scopes)
}

func TestScopesRootItself(t *testing.T) {
	root := scopeTree(t)
	scopes, e := lint.Scopes(root, root, []string{"."})
	assert.Nil(t, e)
	assert.Strings(t, nil, scopes)
}

func TestScopesMissing(t *testing.T) {
	root := scopeTree(t)
	_, e := lint.Scopes(root, root, []string{"doc/ai/spec/ghost.md"})
	assert.Error(t, e)
}

func TestScopesOutside(t *testing.T) {
	root := scopeTree(t)
	_, e := lint.Scopes(root, root, []string{".."})
	assert.Error(t, e)
}

func scopeTree(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	testutil.WriteFile(t, root, "doc/ai/spec/naming.md", "# Naming\n")
	testutil.WriteFile(t, root, "pkg/lint/lint.go", "package lint\n")

	return root
}
