package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestRenameRetargetsSymbolAndItsFile(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("rename-follow/src"))
	system_tester.Write(
		t,
		o.Directory,
		"doc/note.md",
		"Call `go:pkg/alfa.OldName` from `pkg/alfa/old_name.go`.\n",
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"OldName",
		"NewName",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"Call `go:pkg/alfa.NewName` from `pkg/alfa/new_name.go`.\n",
		o.Read("doc/note.md"),
	)
}

func TestRenameMethodRetargetsQualifiedName(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("rename-method/src"))
	system_tester.Write(
		t,
		o.Directory,
		"doc/note.md",
		"Look up with `go:pkg/target.Store.FindByName()`.\n",
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"FindByName",
		"LookupByName",
		"Store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"Look up with `go:pkg/target.Store.LookupByName()`.\n",
		o.Read("doc/note.md"),
	)
}

func TestMoveRetargetsExportedSymbol(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-visibility/src"),
	)
	system_tester.Write(
		t,
		o.Directory,
		"doc/note.md",
		"See `go:pkg/alfa.lone`.\n",
	)
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		[]string{"lone"},
		"",
		"example/pkg/bravo",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(t, "See `go:pkg/bravo.Lone`.\n", o.Read("doc/note.md"))
}

func TestMoveMethodsRetargetsMethodSpan(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-methods/src"))
	system_tester.Write(
		t,
		o.Directory,
		"doc/note.md",
		"See `go:pkg/alfa.Service.Greet`.\n",
	)
	r, e := o.Symbol.MoveMethods(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"Service",
		[]string{"Greet"},
		"example/pkg/bravo",
		"Worker",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"See `go:pkg/bravo.Worker.Greet`.\n",
		o.Read("doc/note.md"),
	)
}

func TestMoveMethodsRetargetsExportedBackReference(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-methods/src"))
	system_tester.Write(
		t,
		o.Directory,
		"doc/note.md",
		"Uses `go:pkg/alfa.prefix`.\n",
	)
	r, e := o.Symbol.MoveMethods(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"Service",
		[]string{"Label"},
		"example/pkg/bravo",
		"Worker",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(t, "Uses `go:pkg/alfa.Prefix`.\n", o.Read("doc/note.md"))
}

func TestMoveSymbolsRetargetsExportedBackReference(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("back-reference-unexported/src"),
	)
	system_tester.Write(
		t,
		o.Directory,
		"doc/note.md",
		"Uses `go:pkg/source.stayer`.\n",
	)
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/source",
		[]string{"Mover"},
		"",
		"example/pkg/target",
		"mover.go",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(t, "Uses `go:pkg/source.Stayer`.\n", o.Read("doc/note.md"))
}

func TestExtractTypeRetargetsExportedBackReference(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-dependency/src"),
	)
	system_tester.Write(
		t,
		o.Directory,
		"doc/note.md",
		"Uses `go:pkg/target.prefix`.\n",
	)
	r, e := o.Symbol.ExtractType(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"Store",
		"example/pkg/model",
		"",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(t, "Uses `go:pkg/target.Prefix`.\n", o.Read("doc/note.md"))
}
