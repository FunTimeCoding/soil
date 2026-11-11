package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"strings"
	"testing"
)

func TestRenameEmbeddedTypeRenamesItsField(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-embedded/src"),
	)
	r, e := service_tester.NewService(t).Symbol.Rename(
		t.Context(),
		d,
		"example/pkg/alfa",
		"Reserve",
		"Budget",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	reported := 0

	for _, c := range r.Entries {
		if strings.HasSuffix(c.Path, "pkg/bravo/holder.go") {
			reported++
		}
	}

	assert.Integer(t, 4, reported)
	assert.String(
		t,
		"package bravo\n\nimport \"example/pkg/alfa\"\n\ntype Holder struct {\n\talfa.Budget\n\tName string\n}\n\nfunc New() *Holder {\n\treturn &Holder{Budget: alfa.Budget{Amount: 1}, Name: \"bravo\"}\n}\n\nfunc (h *Holder) Total() int {\n\treturn h.Budget.Amount\n}\n",
		service_tester.ReadFixtureFile(t, d, "pkg/bravo/holder.go"),
	)
	assert.String(
		t,
		"package charlie\n\nimport \"example/pkg/alfa\"\n\ntype Reserve = alfa.Budget\n\ntype Keeper struct {\n\tReserve\n}\n\nfunc Make() Keeper {\n\treturn Keeper{Reserve: Reserve{Amount: 2}}\n}\n",
		service_tester.ReadFixtureFile(t, d, "pkg/charlie/keeper.go"),
	)
	assertCompiles(t, d)
}

func TestRenameEmbeddedTypeRefusesCapture(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-embedded-capture/src"),
	)
	r, e := service_tester.NewService(t).Symbol.Rename(
		t.Context(),
		d,
		"example/pkg/alfa",
		"Reserve",
		"Budget",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "Holder already reaches Budget")
	assert.StringContains(
		t,
		"\talfa.Reserve\n",
		service_tester.ReadFixtureFile(t, d, "pkg/bravo/holder.go"),
	)
}

func TestMoveExportingEmbeddedTypeRenamesItsField(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-embedded-flip/src"),
	)
	r, e := service_tester.NewService(t).Symbol.MoveSymbol(
		t.Context(),
		d,
		"example/pkg/alfa",
		"reserve",
		"example/pkg/bravo",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := service_tester.ReadFixtureFile(t, d, "pkg/alfa/reserve.go")
	assert.StringContains(
		t,
		"Holder{Reserve: bravo.Reserve{Amount: 1}}",
		source,
	)
	assert.StringContains(t, "h.Reserve.Amount", source)
	assertCompiles(t, d)
}

func TestExtractTypeExportingEmbeddedTypeRenamesItsField(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-embedded-flip/src"),
	)
	r, e := service_tester.NewService(t).Symbol.ExtractType(
		t.Context(),
		d,
		"example/pkg/alfa",
		"reserve",
		"example/pkg/bravo",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := service_tester.ReadFixtureFile(t, d, "pkg/alfa/reserve.go")
	assert.StringContains(
		t,
		"Holder{Reserve: bravo.Reserve{Amount: 1}}",
		source,
	)
	assertCompiles(t, d)
}

func TestBackReferenceExportingEmbeddedTypeRenamesItsField(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-embedded-flip/src"),
	)
	r, e := service_tester.NewService(t).Symbol.MoveSymbol(
		t.Context(),
		d,
		"example/pkg/alfa",
		"Mover",
		"example/pkg/bravo",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := service_tester.ReadFixtureFile(t, d, "pkg/alfa/reserve.go")
	assert.StringContains(t, "Holder{Reserve: Reserve{Amount: 1}}", source)
	assert.StringContains(t, "h.Reserve.Amount", source)
	assertCompiles(t, d)
}
