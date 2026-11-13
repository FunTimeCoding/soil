package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/pointer_cache"
	"github.com/funtimecoding/soil/pkg/source/warm_forest"
	"testing"
)

func TestPointerCacheSeesCreatedFile(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(
		t,
		root,
		"doc/alfa.md",
		"# Alfa\n\nSee `doc/bravo.md`.\n",
	)
	f := warm_forest.New()
	c := pointer_cache.New()
	assert.Count(t, 1, resolvedRun(t, root, f, c).Entries)
	assert.Count(t, 1, resolvedRun(t, root, f, c).Entries)
	testutil.WriteFile(t, root, "doc/bravo.md", "# Bravo\n")
	assert.Count(t, 0, resolvedRun(t, root, f, c).Entries)
}

func TestPointerCacheSeesEditedHeading(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(
		t,
		root,
		"doc/alfa.md",
		"# Alfa\n\nSee `doc/bravo.md#charlie`.\n",
	)
	testutil.WriteFile(t, root, "doc/bravo.md", "# Bravo\n")
	f := warm_forest.New()
	c := pointer_cache.New()
	assert.Count(t, 1, resolvedRun(t, root, f, c).Entries)
	testutil.WriteFile(t, root, "doc/bravo.md", "# Bravo\n\n## Charlie\n")
	assert.Count(t, 0, resolvedRun(t, root, f, c).Entries)
}

func TestPointerCacheReplaysUncheckedUntilTestExists(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(t, root, "doc/alfa.md", "# Alfa\n\nSee `TestBravo`.\n")
	testutil.WriteFile(t, root, "pkg/echo/echo.go", "package echo\n")
	f := warm_forest.New()
	c := pointer_cache.New()
	assert.Count(t, 1, resolvedRun(t, root, f, c).Unchecked)
	assert.Count(t, 1, resolvedRun(t, root, f, c).Unchecked)
	testutil.WriteFile(
		t,
		root,
		"pkg/charlie/unit/charlie_test.go",
		"package unit\n\nimport \"testing\"\n\nfunc TestBravo(t *testing.T) {}\n",
	)
	assert.Count(t, 0, resolvedRun(t, root, f, c).Unchecked)
}

func TestPointerCacheSeesCodeLiteral(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(
		t,
		root,
		"doc/alfa.md",
		"# Alfa\n\nSee `ghcr.io/bravo/charlie`.\n",
	)
	testutil.WriteFile(t, root, "pkg/echo/echo.go", "package echo\n")
	f := warm_forest.New()
	c := pointer_cache.New()
	assert.Count(t, 1, resolvedRun(t, root, f, c).Unchecked)
	testutil.WriteFile(
		t,
		root,
		"pkg/delta/delta.go",
		"package delta\n\nconst Image = \"ghcr.io/bravo/charlie\"\n",
	)
	assert.Count(t, 0, resolvedRun(t, root, f, c).Unchecked)
}

func TestPointerCacheSeesNewDeclaration(t *testing.T) {
	root := t.TempDir()
	testutil.WriteModFile(t, root, "example.org/alfa")
	testutil.WriteFile(
		t,
		root,
		"doc/bravo.md",
		"# Bravo\n\nSee `go:pkg/charlie.Delta`.\n",
	)
	testutil.WriteFile(t, root, "pkg/charlie/charlie.go", "package charlie\n")
	f := warm_forest.New()
	c := pointer_cache.New()
	assert.Count(t, 1, resolvedRun(t, root, f, c).Entries)
	testutil.WriteFile(
		t,
		root,
		"pkg/charlie/charlie.go",
		"package charlie\n\nfunc Delta() {}\n",
	)
	assert.Count(t, 0, resolvedRun(t, root, f, c).Entries)
}

func TestPointerCacheSeesRouteSpecification(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(
		t,
		root,
		"doc/alfa.md",
		"---\nbase: pkg/bravo\n---\n\n# Alfa\n\nSee `route:/api/charlie`.\n",
	)
	testutil.WriteFile(t, root, "pkg/bravo/bravo.go", "package bravo\n")
	f := warm_forest.New()
	c := pointer_cache.New()
	assert.Count(t, 1, resolvedRun(t, root, f, c).Entries)
	testutil.WriteFile(
		t,
		root,
		"pkg/bravo/generated/server/openapi.yaml",
		"---\npaths:\n  /api/charlie:\n    get: {}\n",
	)
	assert.Count(t, 0, resolvedRun(t, root, f, c).Entries)
}

func TestPointerCacheSeesRouteRemoved(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(
		t,
		root,
		"doc/alfa.md",
		"---\nbase: pkg/bravo\n---\n\n# Alfa\n\nSee `route:/api/charlie`.\n",
	)
	testutil.WriteFile(
		t,
		root,
		"pkg/bravo/generated/server/openapi.yaml",
		"---\npaths:\n  /api/charlie:\n    get: {}\n",
	)
	f := warm_forest.New()
	c := pointer_cache.New()
	assert.Count(t, 0, resolvedRun(t, root, f, c).Entries)
	testutil.WriteFile(
		t,
		root,
		"pkg/bravo/generated/server/openapi.yaml",
		"---\npaths:\n  /api/delta:\n    get: {}\n",
	)
	assert.Count(t, 1, resolvedRun(t, root, f, c).Entries)
}

func TestPointerCacheSeesIgnoreFile(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(
		t,
		root,
		"doc/alfa.md",
		"# Alfa\n\nSee `doc/bravo.txt`.\n",
	)
	f := warm_forest.New()
	c := pointer_cache.New()
	assert.Count(t, 1, resolvedRun(t, root, f, c).Entries)
	testutil.WriteFile(t, root, ".gitignore", "doc/bravo.txt\n")
	assert.Count(t, 0, resolvedRun(t, root, f, c).Entries)
}

func TestPointerCacheAuditCatchesHiddenChange(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(
		t,
		root,
		"doc/alfa.md",
		"# Alfa\n\nSee `doc/bravo.md`.\n",
	)
	f := warm_forest.New()
	c := pointer_cache.New()
	var missed []string
	c.Mismatch = func(
		_ string,
		document string,
		_ string,
		_ string,
	) {
		missed = append(missed, document)
	}
	assert.Count(t, 1, resolvedRun(t, root, f, c).Entries)
	testutil.WriteFile(t, root, "doc/bravo.md", "# Bravo\n")
	f.Refresh()

	for range 17 {
		c.Begin(map[string][]string{})
	}

	assert.Count(t, 1, resolvedRun(t, root, f, c).Entries)
	assert.Strings(t, nil, missed)
	assert.Count(t, 0, resolvedRun(t, root, f, c).Entries)
	assert.Strings(t, []string{"doc/alfa.md"}, missed)
}

func TestPointerCacheKeepsNothingAskedOutsideTrees(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(
		t,
		root,
		"doc/alfa.md",
		"# Alfa\n\nSee `../bravo/charlie.md`.\n",
	)
	f := warm_forest.New()
	c := pointer_cache.New()
	assert.Count(t, 1, resolvedRun(t, root, f, c).Entries)
	testutil.WriteFile(t, root, "../bravo/charlie.md", "# Charlie\n")
	assert.Count(t, 0, resolvedRun(t, root, f, c).Entries)
}
