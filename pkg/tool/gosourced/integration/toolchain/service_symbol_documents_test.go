package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestRenameRetargetsSymbolAndItsFile(t *testing.T) {
	d := renameFollowFixture(t)
	testutil.WriteFile(
		t,
		d,
		"doc/note.md",
		"Call `go:pkg/alfa.OldName` from `pkg/alfa/old_name.go`.\n",
	)
	r, e := service_tester.NewService(t).Symbol.Rename(
		t.Context(),
		d,
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
		service_tester.ReadFixtureFile(t, d, "doc/note.md"),
	)
}

func TestRenameMethodRetargetsQualifiedName(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-method/src"),
	)
	testutil.WriteFile(
		t,
		d,
		"doc/note.md",
		"Look up with `go:pkg/target.Store.FindByName()`.\n",
	)
	r, e := service_tester.NewService(t).Symbol.Rename(
		t.Context(),
		d,
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
		service_tester.ReadFixtureFile(t, d, "doc/note.md"),
	)
}

func TestMoveRetargetsExportedSymbol(t *testing.T) {
	d := moveVisibilityFixture(t)
	testutil.WriteFile(t, d, "doc/note.md", "See `go:pkg/alfa.lone`.\n")
	r, e := service_tester.NewService(t).Symbol.MoveSymbols(
		t.Context(),
		d,
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
	assert.String(
		t,
		"See `go:pkg/bravo.Lone`.\n",
		service_tester.ReadFixtureFile(t, d, "doc/note.md"),
	)
}

func TestMoveMethodsRetargetsMethodSpan(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-methods/src"),
	)
	testutil.WriteFile(
		t,
		d,
		"doc/note.md",
		"See `go:pkg/alfa.Service.Greet`.\n",
	)
	r, e := service_tester.NewService(t).Symbol.MoveMethods(
		t.Context(),
		d,
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
		service_tester.ReadFixtureFile(t, d, "doc/note.md"),
	)
}

func TestMoveMethodsRetargetsExportedBackReference(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-methods/src"),
	)
	testutil.WriteFile(t, d, "doc/note.md", "Uses `go:pkg/alfa.prefix`.\n")
	r, e := service_tester.NewService(t).Symbol.MoveMethods(
		t.Context(),
		d,
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
	assert.String(
		t,
		"Uses `go:pkg/alfa.Prefix`.\n",
		service_tester.ReadFixtureFile(t, d, "doc/note.md"),
	)
}

func TestMoveSymbolsRetargetsExportedBackReference(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("back-reference-unexported/src"),
	)
	testutil.WriteFile(t, d, "doc/note.md", "Uses `go:pkg/source.stayer`.\n")
	r, e := service_tester.NewService(t).Symbol.MoveSymbols(
		t.Context(),
		d,
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
	assert.String(
		t,
		"Uses `go:pkg/source.Stayer`.\n",
		service_tester.ReadFixtureFile(t, d, "doc/note.md"),
	)
}

func TestExtractTypeRetargetsExportedBackReference(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-dependency/src"),
	)
	testutil.WriteFile(t, d, "doc/note.md", "Uses `go:pkg/target.prefix`.\n")
	r, e := service_tester.NewService(t).Symbol.ExtractType(
		t.Context(),
		d,
		"example/pkg/target",
		"Store",
		"example/pkg/model",
		"",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"Uses `go:pkg/target.Prefix`.\n",
		service_tester.ReadFixtureFile(t, d, "doc/note.md"),
	)
}
