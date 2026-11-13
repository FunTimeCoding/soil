package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"golang.org/x/tools/go/packages"
	"path/filepath"
	"testing"
)

func TestIgnoreStubOverlayHidesMovedFile(t *testing.T) {
	root, e := filepath.EvalSymlinks(t.TempDir())
	assert.FatalOnError(t, e)
	testutil.WriteModFile(t, root, "example")
	testutil.WriteFile(
		t,
		root,
		"alfa/alfa.go",
		"package alfa\n\nfunc Bravo() {}\n",
	)
	testutil.WriteFile(
		t,
		root,
		"alfa/charlie.go",
		"package alfa\n\nfunc Charlie() {}\n",
	)
	testutil.WriteFile(t, root, "delta/delta.go", "package delta\n")
	testutil.WriteFile(
		t,
		root,
		"echo/echo.go",
		"package echo\n\nimport \"example/alfa\"\n\nfunc Foxtrot() { alfa.Charlie() }\n",
	)
	overlay := map[string][]byte{
		filepath.Join(root, "alfa/charlie.go"): []byte(
			"//go:build ignore\n\npackage alfa\n",
		),
		filepath.Join(root, "delta/charlie.go"): []byte(
			"package delta\n\nfunc Charlie() {}\n",
		),
		filepath.Join(root, "echo/echo.go"): []byte(
			"package echo\n\nimport \"example/delta\"\n\nfunc Foxtrot() { delta.Charlie() }\n",
		),
	}
	loaded, _, f := resolve.LoadPackagesUnchecked(root, overlay, "./...")
	assert.FatalOnError(t, f)
	byPath := map[string]*packages.Package{}

	for _, p := range loaded {
		assert.Count(t, 0, p.Errors)
		byPath[p.PkgPath] = p
	}

	assert.True(
		t,
		byPath["example/alfa"].Types.Scope().Lookup("Charlie") == nil,
	)
	assert.True(
		t,
		byPath["example/delta"].Types.Scope().Lookup("Charlie") != nil,
	)
}

func TestIgnoreStubOverlayOnWholePackage(t *testing.T) {
	root, e := filepath.EvalSymlinks(t.TempDir())
	assert.FatalOnError(t, e)
	testutil.WriteModFile(t, root, "example")
	testutil.WriteFile(
		t,
		root,
		"alfa/alfa.go",
		"package alfa\n\nfunc Bravo() {}\n",
	)
	testutil.WriteFile(t, root, "delta/delta.go", "package delta\n")
	overlay := map[string][]byte{
		filepath.Join(root, "alfa/alfa.go"): []byte(
			"//go:build ignore\n\npackage alfa\n",
		),
	}
	loaded, _, f := resolve.LoadPackagesUnchecked(root, overlay, "./...")
	assert.FatalOnError(t, f)
	var paths []string

	for _, p := range loaded {
		assert.Count(t, 0, p.Errors)
		paths = append(paths, p.PkgPath)
	}

	assert.Strings(t, []string{"example/delta"}, paths)
}
