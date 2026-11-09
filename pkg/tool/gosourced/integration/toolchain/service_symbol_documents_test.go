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
