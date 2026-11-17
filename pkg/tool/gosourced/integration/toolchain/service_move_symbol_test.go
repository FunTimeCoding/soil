package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMoveConstant(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-constant/src"))
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"itemFields",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	_, e = os.Stat(filepath.Join(o.Directory, "pkg/target/constant.go"))
	assert.True(t, os.IsNotExist(e))
	moved := o.Read("pkg/target/constant/constant.go")
	rewrite_tester.Formatted(t, moved)
	assert.StringContains(t, "ItemFields = \"alfa\"", moved)
	assert.StringContains(t, "Host = \"bravo\"", moved)
	run := o.Read("pkg/target/run.go")
	rewrite_tester.Formatted(t, run)
	assert.StringContains(t, "return constant.ItemFields", run)
	assert.StringContains(t, "example/pkg/target/constant", run)
}

func TestMoveIntoReferenced(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-into-referenced/src"),
	)
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"Table",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := o.Read("pkg/target/constant/constant.go")
	rewrite_tester.Formatted(t, moved)
	assert.StringContains(t, "Address: \"primary\"", moved)
	assert.StringNotContains(t, "example/pkg/target/constant", moved)
	run := o.Read("pkg/target/run.go")
	rewrite_tester.Formatted(t, run)
	assert.StringContains(t, "return constant.Table", run)
	assert.StringContains(t, "example/pkg/target/constant", run)
}

func TestMoveCrossPackage(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-cross-package/src"),
	)
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"Mode",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	run := o.Read("pkg/target/run.go")
	rewrite_tester.Formatted(t, run)
	assert.StringContains(t, "return constant.Mode", run)
	caller := o.Read("pkg/caller/run.go")
	rewrite_tester.Formatted(t, caller)
	assert.StringContains(t, "return targetConstant.Mode", caller)
	assert.StringContains(t, "example/pkg/target/constant", caller)
	assert.False(t, strings.Contains(caller, "target.Mode"))
}

func TestMoveAlias(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-alias/src"))
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"Mode",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	caller := o.Read("pkg/caller/run.go")
	rewrite_tester.Formatted(t, caller)
	assert.StringContains(t, "targetConstant.Mode", caller)
	assert.StringContains(
		t,
		"targetConstant \"example/pkg/target/constant\"",
		caller,
	)
	assert.StringContains(t, "constant.Local", caller)
}

func TestMoveCreateRefused(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-create/src"))
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"itemFields",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestMoveCreate(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-create/src"))
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"itemFields",
		"example/pkg/target/constant",
		"constant.go",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := o.Read("pkg/target/constant/constant.go")
	rewrite_tester.Formatted(t, moved)
	assert.StringContains(t, "package constant", moved)
	assert.StringContains(t, "ItemFields = \"alfa\"", moved)
	run := o.Read("pkg/target/run.go")
	rewrite_tester.Formatted(t, run)
	assert.StringContains(t, "return constant.ItemFields", run)
}

func TestMoveCollision(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-collision/src"))
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"itemFields",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestMoveCycle(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-cycle/src"))
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/target/constant",
		"Mode",
		"example/pkg/other",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "cycle")
}

func TestMoveIntoImporterOnlyOfMovedSymbol(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-into-importer/src"),
	)
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/target/constant",
		"Mode",
		"example/pkg/other",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"package other\n\nfunc Run() string {\n\treturn Mode\n}\n",
		o.Read("pkg/other/run.go"),
	)
}

func TestMoveIntoImporterOfSourceWithoutCycle(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-beside-source/src"),
	)
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
		"package constant\n\nconst Kind = \"kind\"\n",
		o.Read("pkg/alfa/constant/constant.go"),
	)
	assert.StringContains(
		t,
		"const Mode = \"alfa\"",
		o.Read("pkg/bravo/constant/constant.go"),
	)
}

func TestMoveCycleThroughImporterLosingSource(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-cycle-through-importer/src"),
	)
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/source",
		"Mode",
		"example/pkg/target",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "example/pkg/relay")
}

func TestMoveCycleThroughCarriedImport(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-cycle-through-carried/src"),
	)
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/source",
		"Mode",
		"example/pkg/target",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(
		t,
		r,
		"example/pkg/target → example/pkg/relay → example/pkg/target",
	)
}

func TestMoveSymbolQualifiesSibling(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-dependency/src"),
	)
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"composed",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	rewrite_tester.FixedContains(t, "prefix → Prefix", r)
	assert.StringContains(
		t,
		"const composed = target.Prefix + \"bravo\"",
		o.Read("pkg/target/constant/constant.go"),
	)
	assert.StringContains(
		t,
		"const Prefix = \"alfa\"",
		o.Read("pkg/target/constant.go"),
	)
	rewrite_tester.Compiles(t, o.Directory)
}

func TestMoveGroup(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-group/src"))
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"First",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := o.Read("pkg/target/constant.go")
	rewrite_tester.Formatted(t, source)
	assert.StringContains(t, "Second = \"bravo\"", source)
	assert.False(t, strings.Contains(source, "First"))
	moved := o.Read("pkg/target/constant/constant.go")
	rewrite_tester.Formatted(t, moved)
	assert.StringContains(t, "First = \"alfa\"", moved)
	run := o.Read("pkg/target/run.go")
	rewrite_tester.Formatted(t, run)
	assert.StringContains(t, "constant.First + Second", run)
}

func TestMoveFunction(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-function/src"))
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"IsValid",
		"example/pkg/check",
		"",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	_, e = os.Stat(filepath.Join(o.Directory, "pkg/target/is_valid.go"))
	assert.True(t, os.IsNotExist(e))
	moved := o.Read("pkg/check/is_valid.go")
	rewrite_tester.Formatted(t, moved)
	assert.StringContains(t, "package check", moved)
	assert.StringContains(t, "func IsValid(", moved)
	caller := o.Read("pkg/caller/run.go")
	rewrite_tester.Formatted(t, caller)
	assert.StringContains(t, "check.IsValid(\"alfa\")", caller)
	assert.StringContains(t, "example/pkg/check", caller)
}
