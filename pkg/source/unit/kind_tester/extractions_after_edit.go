package kind_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
)

func (o *Tester) ExtractionsAfterEdit(edited string) int {
	o.t.Helper()
	directory := o.t.TempDir()
	testutil.WriteModFile(o.t, directory, "example")
	system_tester.New(o.t).Write(
		directory,
		"alfa/alfa.go",
		"package alfa\n\ntype Alfa struct{}\n\ntype echo struct{}\n",
	)
	system_tester.New(o.t).Write(
		directory,
		"bravo/bravo.go",
		"package bravo\n\nimport alfa \"example/alfa\"\n\nfunc Bravo() alfa.Alfa {\n\treturn alfa.Alfa{}\n}\n",
	)
	store := o.t.TempDir()
	var extracted, external int
	counting := Counting("golf", &extracted, &external)
	index.Facts[*string](index.New(store, directory, counting), counting)
	system_tester.New(o.t).Write(directory, "alfa/alfa.go", edited)
	index.Facts[*string](index.New(store, directory, counting), counting)

	return extracted
}
