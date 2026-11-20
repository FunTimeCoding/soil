package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/module_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestMoveIntoOwnConstantDropsOrphan(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-orphan/src"))
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
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
		o.Read("pkg/bravo/run.go"),
	)
}

func TestMoveIntoForeignConstantCarriesItsDomain(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-foreign-constant/src"),
	)
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"Mode",
		"example/pkg/delta/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	caller := o.Read("pkg/charlie/run.go")
	assert.StringContains(
		t,
		"deltaConstant \"example/pkg/delta/constant\"",
		caller,
	)
	assert.StringContains(t, "fmt.Sprint(deltaConstant.Mode)", caller)
}

func TestMoveIntoNewFileNamesCarriedConstants(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-two-constants/src"),
	)
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/xray/bravo",
		"Clamp",
		"example/pkg/xray/charlie",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"package charlie\n\nimport (\n\talfaConstant \"example/pkg/alfa/constant\"\n\t\"example/pkg/xray/constant\"\n)\n\nfunc Clamp(name string) int {\n\tif name == alfaConstant.Prefix {\n\t\treturn constant.Limit\n\t}\n\n\treturn 0\n}\n",
		o.Read("pkg/xray/charlie/clamp.go"),
	)
}

func TestMoveIntoExistingFileNamesCarriedConstants(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-two-constants/src"),
	)
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/xray/bravo",
		"Clamp",
		"example/pkg/xray/charlie",
		"size.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"package charlie\n\nimport (\n\talfaConstant \"example/pkg/alfa/constant\"\n\t\"example/pkg/xray/constant\"\n)\n\nfunc Size() int {\n\treturn constant.Limit * 2\n}\n\nfunc Clamp(name string) int {\n\tif name == alfaConstant.Prefix {\n\t\treturn constant.Limit\n\t}\n\n\treturn 0\n}\n",
		o.Read("pkg/xray/charlie/size.go"),
	)
}

func TestMoveKeepsOneImportBlock(t *testing.T) {
	o := module_tester.New(t).DottedModule("move-dotted/src")
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
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
		o.Read("pkg/charlie/run.go"),
	)
}
