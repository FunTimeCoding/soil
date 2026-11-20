package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"strings"
	"testing"
)

func TestRenameEmbeddedTypeRenamesItsField(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-embedded/src"),
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
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
		o.Read("pkg/bravo/holder.go"),
	)
	assert.String(
		t,
		"package charlie\n\nimport \"example/pkg/alfa\"\n\ntype Reserve = alfa.Budget\n\ntype Keeper struct {\n\tReserve\n}\n\nfunc Make() Keeper {\n\treturn Keeper{Reserve: Reserve{Amount: 2}}\n}\n",
		o.Read("pkg/charlie/keeper.go"),
	)
	rewrite_tester.New(t).Compiles(o.Directory)
}

func TestRenameEmbeddedTypeRefusesCapture(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-embedded-capture/src"),
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"Reserve",
		"Budget",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "Holder already reaches Budget")
	assert.StringContains(t, "\talfa.Reserve\n", o.Read("pkg/bravo/holder.go"))
}

func TestMoveExportingEmbeddedTypeRenamesItsField(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-embedded-flip/src"),
	)
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"reserve",
		"example/pkg/bravo",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := o.Read("pkg/alfa/reserve.go")
	assert.StringContains(
		t,
		"Holder{Reserve: bravo.Reserve{Amount: 1}}",
		source,
	)
	assert.StringContains(t, "h.Reserve.Amount", source)
	rewrite_tester.New(t).Compiles(o.Directory)
}

func TestExtractTypeExportingEmbeddedTypeRenamesItsField(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-embedded-flip/src"),
	)
	r, e := o.Symbol.ExtractType(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"reserve",
		"example/pkg/bravo",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := o.Read("pkg/alfa/reserve.go")
	assert.StringContains(
		t,
		"Holder{Reserve: bravo.Reserve{Amount: 1}}",
		source,
	)
	rewrite_tester.New(t).Compiles(o.Directory)
}

func TestBackReferenceExportingEmbeddedTypeRenamesItsField(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-embedded-flip/src"),
	)
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"Mover",
		"example/pkg/bravo",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := o.Read("pkg/alfa/reserve.go")
	assert.StringContains(t, "Holder{Reserve: Reserve{Amount: 1}}", source)
	assert.StringContains(t, "h.Reserve.Amount", source)
	rewrite_tester.New(t).Compiles(o.Directory)
}
