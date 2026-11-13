package toolchain

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/unit/kind_tester"
	"testing"
)

func extractionsAfterEdit(
	t *testing.T,
	edited string,
) int {
	t.Helper()
	directory := t.TempDir()
	testutil.WriteModFile(t, directory, "example")
	testutil.WriteFile(
		t,
		directory,
		"alfa/alfa.go",
		"package alfa\n\ntype Alfa struct{}\n\ntype echo struct{}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"bravo/bravo.go",
		"package bravo\n\nimport alfa \"example/alfa\"\n\nfunc Bravo() alfa.Alfa {\n\treturn alfa.Alfa{}\n}\n",
	)
	store := t.TempDir()
	var extracted, external int
	counting := kind_tester.Counting("golf", &extracted, &external)
	index.Facts[*string](index.New(store, directory, counting), counting)
	testutil.WriteFile(t, directory, "alfa/alfa.go", edited)
	index.Facts[*string](index.New(store, directory, counting), counting)

	return extracted
}
