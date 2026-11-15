package kind_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"testing"
)

func ExtractionsAfterEdit(
	t *testing.T,
	edited string,
) int {
	t.Helper()
	directory := t.TempDir()
	testutil.WriteModFile(t, directory, "example")
	system_tester.Write(
		t,
		directory,
		"alfa/alfa.go",
		"package alfa\n\ntype Alfa struct{}\n\ntype echo struct{}\n",
	)
	system_tester.Write(
		t,
		directory,
		"bravo/bravo.go",
		"package bravo\n\nimport alfa \"example/alfa\"\n\nfunc Bravo() alfa.Alfa {\n\treturn alfa.Alfa{}\n}\n",
	)
	store := t.TempDir()
	var extracted, external int
	counting := Counting("golf", &extracted, &external)
	index.Facts[*string](index.New(store, directory, counting), counting)
	system_tester.Write(t, directory, "alfa/alfa.go", edited)
	index.Facts[*string](index.New(store, directory, counting), counting)

	return extracted
}
