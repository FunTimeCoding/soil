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
	d := service_tester.Prepare(t, "move-constant/src")
	s := service_tester.NewService(t)
	r, e := s.Symbol.MoveSymbol(
		t.Context(),
		d,
		"example/pkg/target",
		"itemFields",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	_, e = os.Stat(filepath.Join(d, "pkg/target/constant.go"))
	assert.True(t, os.IsNotExist(e))
	moved := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/constant/constant.go",
	)
	rewrite_tester.Formatted(t, moved)
	assert.StringContains(t, "ItemFields = \"alfa\"", moved)
	assert.StringContains(t, "Host = \"bravo\"", moved)
	run := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	rewrite_tester.Formatted(t, run)
	assert.StringContains(t, "return constant.ItemFields", run)
	assert.StringContains(t, "example/pkg/target/constant", run)
}

func TestMoveIntoReferenced(t *testing.T) {
	d := service_tester.Prepare(t, "move-into-referenced/src")
	s := service_tester.NewService(t)
	r, e := s.Symbol.MoveSymbol(
		t.Context(),
		d,
		"example/pkg/target",
		"Table",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/constant/constant.go",
	)
	rewrite_tester.Formatted(t, moved)
	assert.StringContains(t, "Address: \"primary\"", moved)
	assert.StringNotContains(t, "example/pkg/target/constant", moved)
	run := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	rewrite_tester.Formatted(t, run)
	assert.StringContains(t, "return constant.Table", run)
	assert.StringContains(t, "example/pkg/target/constant", run)
}

func TestMoveCrossPackage(t *testing.T) {
	d := service_tester.Prepare(t, "move-cross-package/src")
	s := service_tester.NewService(t)
	r, e := s.Symbol.MoveSymbol(
		t.Context(),
		d,
		"example/pkg/target",
		"Mode",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	run := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	rewrite_tester.Formatted(t, run)
	assert.StringContains(t, "return constant.Mode", run)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go")
	rewrite_tester.Formatted(t, caller)
	assert.StringContains(t, "return targetConstant.Mode", caller)
	assert.StringContains(t, "example/pkg/target/constant", caller)
	assert.False(t, strings.Contains(caller, "target.Mode"))
}

func TestMoveAlias(t *testing.T) {
	d := service_tester.Prepare(t, "move-alias/src")
	s := service_tester.NewService(t)
	r, e := s.Symbol.MoveSymbol(
		t.Context(),
		d,
		"example/pkg/target",
		"Mode",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go")
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
	d := service_tester.Prepare(t, "move-create/src")
	s := service_tester.NewService(t)
	r, e := s.Symbol.MoveSymbol(
		t.Context(),
		d,
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
	d := service_tester.Prepare(t, "move-create/src")
	s := service_tester.NewService(t)
	r, e := s.Symbol.MoveSymbol(
		t.Context(),
		d,
		"example/pkg/target",
		"itemFields",
		"example/pkg/target/constant",
		"constant.go",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/constant/constant.go",
	)
	rewrite_tester.Formatted(t, moved)
	assert.StringContains(t, "package constant", moved)
	assert.StringContains(t, "ItemFields = \"alfa\"", moved)
	run := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	rewrite_tester.Formatted(t, run)
	assert.StringContains(t, "return constant.ItemFields", run)
}

func TestMoveCollision(t *testing.T) {
	d := service_tester.Prepare(t, "move-collision/src")
	s := service_tester.NewService(t)
	r, e := s.Symbol.MoveSymbol(
		t.Context(),
		d,
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
	d := service_tester.Prepare(t, "move-cycle/src")
	s := service_tester.NewService(t)
	r, e := s.Symbol.MoveSymbol(
		t.Context(),
		d,
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
	d := service_tester.Prepare(t, "move-into-importer/src")
	r, e := service_tester.NewService(t).Symbol.MoveSymbol(
		t.Context(),
		d,
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
		service_tester.ReadFixtureFile(t, d, "pkg/other/run.go"),
	)
}

func TestMoveIntoImporterOfSourceWithoutCycle(t *testing.T) {
	d := service_tester.Prepare(t, "move-beside-source/src")
	r, e := service_tester.NewService(t).Symbol.MoveSymbol(
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
		"package constant\n\nconst Kind = \"kind\"\n",
		service_tester.ReadFixtureFile(t, d, "pkg/alfa/constant/constant.go"),
	)
	assert.StringContains(
		t,
		"const Mode = \"alfa\"",
		service_tester.ReadFixtureFile(t, d, "pkg/bravo/constant/constant.go"),
	)
}

func TestMoveCycleThroughImporterLosingSource(t *testing.T) {
	d := service_tester.Prepare(t, "move-cycle-through-importer/src")
	r, e := service_tester.NewService(t).Symbol.MoveSymbol(
		t.Context(),
		d,
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
	d := service_tester.Prepare(t, "move-cycle-through-carried/src")
	r, e := service_tester.NewService(t).Symbol.MoveSymbol(
		t.Context(),
		d,
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
	d := service_tester.Prepare(t, "move-dependency/src")
	s := service_tester.NewService(t)
	r, e := s.Symbol.MoveSymbol(
		t.Context(),
		d,
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
		service_tester.ReadFixtureFile(t, d, "pkg/target/constant/constant.go"),
	)
	assert.StringContains(
		t,
		"const Prefix = \"alfa\"",
		service_tester.ReadFixtureFile(t, d, "pkg/target/constant.go"),
	)
	rewrite_tester.Compiles(t, d)
}

func TestMoveGroup(t *testing.T) {
	d := service_tester.Prepare(t, "move-group/src")
	s := service_tester.NewService(t)
	r, e := s.Symbol.MoveSymbol(
		t.Context(),
		d,
		"example/pkg/target",
		"First",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := service_tester.ReadFixtureFile(t, d, "pkg/target/constant.go")
	rewrite_tester.Formatted(t, source)
	assert.StringContains(t, "Second = \"bravo\"", source)
	assert.False(t, strings.Contains(source, "First"))
	moved := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/constant/constant.go",
	)
	rewrite_tester.Formatted(t, moved)
	assert.StringContains(t, "First = \"alfa\"", moved)
	run := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	rewrite_tester.Formatted(t, run)
	assert.StringContains(t, "constant.First + Second", run)
}

func TestMoveFunction(t *testing.T) {
	d := service_tester.Prepare(t, "move-function/src")
	s := service_tester.NewService(t)
	r, e := s.Symbol.MoveSymbol(
		t.Context(),
		d,
		"example/pkg/target",
		"IsValid",
		"example/pkg/check",
		"",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	_, e = os.Stat(filepath.Join(d, "pkg/target/is_valid.go"))
	assert.True(t, os.IsNotExist(e))
	moved := service_tester.ReadFixtureFile(t, d, "pkg/check/is_valid.go")
	rewrite_tester.Formatted(t, moved)
	assert.StringContains(t, "package check", moved)
	assert.StringContains(t, "func IsValid(", moved)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go")
	rewrite_tester.Formatted(t, caller)
	assert.StringContains(t, "check.IsValid(\"alfa\")", caller)
	assert.StringContains(t, "example/pkg/check", caller)
}
