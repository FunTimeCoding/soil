package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestMoveIntoOwnConstantDropsTheOrphan(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-orphan/src"),
	)
	r, e := service_tester.NewService(t).MoveSymbol(
		t.Context(),
		d,
		"example/pkg/alfa/constant",
		"Mode",
		"example/pkg/bravo/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"package bravo\n\nimport \"example/pkg/bravo/constant\"\n\nfunc Run() string {\n\treturn constant.Mode\n}\n",
		service_tester.ReadFixtureFile(t, d, "pkg/bravo/run.go"),
	)
}

func TestMoveIntoForeignConstantCarriesItsDomain(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-foreign-constant/src"),
	)
	r, e := service_tester.NewService(t).MoveSymbol(
		t.Context(),
		d,
		"example/pkg/target",
		"Mode",
		"example/pkg/delta/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/charlie/run.go")
	assert.StringContains(
		t,
		"deltaConstant \"example/pkg/delta/constant\"",
		caller,
	)
	assert.StringContains(t, "fmt.Sprint(deltaConstant.Mode)", caller)
}

func TestMoveKeepsOneImportBlock(t *testing.T) {
	d := prepareDottedModule(t, "move-dotted/src")
	r, e := service_tester.NewService(t).MoveSymbol(
		t.Context(),
		d,
		"example.org/m/pkg/target",
		"Mode",
		"example.org/m/pkg/delta/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"package charlie\n\nimport (\n\tdeltaConstant \"example.org/m/pkg/delta/constant\"\n\t\"fmt\"\n)\n\nfunc Run() string {\n\treturn fmt.Sprint(deltaConstant.Mode)\n}\n",
		service_tester.ReadFixtureFile(t, d, "pkg/charlie/run.go"),
	)
}
