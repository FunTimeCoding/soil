package resolve

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestLoadPackagesDiscoversBuildTags(t *testing.T) {
	directory := t.TempDir()
	testutil.WriteModFile(t, directory, "testmodule")
	testutil.WriteFile(
		t,
		directory,
		"pkg/example/example.go",
		"package example\n\nfunc Hello() string { return \"hello\" }\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"pkg/example/example_test.go",
		"//go:build ci\n\npackage example\n\nimport \"testing\"\n\nfunc TestHello(t *testing.T) { Hello() }\n",
	)
	result, _, e := LoadPackages(directory, "./...")
	assert.Nil(t, e)
	var found bool

	for _, p := range result {
		for _, file := range p.Syntax {
			path := p.Fset.File(file.Pos()).Name()

			if path != "" && len(file.Imports) > 0 {
				found = true
			}
		}
	}

	assert.True(t, found)
}

func TestLoadPackagesTagOnlyTestPackage(t *testing.T) {
	directory := t.TempDir()
	testutil.WriteModFile(t, directory, "testmodule")
	testutil.WriteFile(
		t,
		directory,
		"pkg/tagged/tagged_test.go",
		"//go:build ci\n\npackage tagged\n\nimport \"testing\"\n\nfunc TestTagged(t *testing.T) {}\n",
	)
	_, _, e := LoadPackages(directory, "./...")
	assert.Nil(t, e)
}
