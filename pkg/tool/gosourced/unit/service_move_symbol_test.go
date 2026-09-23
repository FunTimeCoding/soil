package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMoveConstant(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-constant/src"),
	)
	s := testService()
	r, e := s.MoveSymbol(
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
	assertFormatted(t, moved)
	assert.StringContains(t, "ItemFields = \"alfa\"", moved)
	assert.StringContains(t, "Host = \"bravo\"", moved)
	run := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	assertFormatted(t, run)
	assert.StringContains(t, "return constant.ItemFields", run)
	assert.StringContains(t, "example/pkg/target/constant", run)
}

func TestMoveIntoReferenced(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-into-referenced/src"),
	)
	s := testService()
	r, e := s.MoveSymbol(
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
	assertFormatted(t, moved)
	assert.StringContains(t, "Address: \"primary\"", moved)
	assert.StringNotContains(t, "example/pkg/target/constant", moved)
	run := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	assertFormatted(t, run)
	assert.StringContains(t, "return constant.Table", run)
	assert.StringContains(t, "example/pkg/target/constant", run)
}

func TestMoveCrossPackage(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-cross-package/src"),
	)
	s := testService()
	r, e := s.MoveSymbol(
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
	assertFormatted(t, run)
	assert.StringContains(t, "return constant.Mode", run)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go")
	assertFormatted(t, caller)
	assert.StringContains(t, "return constant.Mode", caller)
	assert.StringContains(t, "example/pkg/target/constant", caller)
	assert.False(t, strings.Contains(caller, "target.Mode"))
}

func TestMoveAlias(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-alias/src"),
	)
	s := testService()
	r, e := s.MoveSymbol(
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
	assertFormatted(t, caller)
	assert.StringContains(t, "targetConstant.Mode", caller)
	assert.StringContains(
		t,
		"targetConstant \"example/pkg/target/constant\"",
		caller,
	)
	assert.StringContains(t, "constant.Local", caller)
}

func TestMoveCreateRefused(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-create/src"),
	)
	s := testService()
	r, e := s.MoveSymbol(
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
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-create/src"),
	)
	s := testService()
	r, e := s.MoveSymbol(
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
	assertFormatted(t, moved)
	assert.StringContains(t, "package constant", moved)
	assert.StringContains(t, "ItemFields = \"alfa\"", moved)
	run := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	assertFormatted(t, run)
	assert.StringContains(t, "return constant.ItemFields", run)
}

func TestMoveCollision(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-collision/src"),
	)
	s := testService()
	r, e := s.MoveSymbol(
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
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-cycle/src"),
	)
	s := testService()
	r, e := s.MoveSymbol(
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

func TestMoveDependency(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-dependency/src"),
	)
	s := testService()
	r, e := s.MoveSymbol(
		d,
		"example/pkg/target",
		"composed",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "package-local")
	testutil.AssertBlockedContains(t, r, "prefix")
}

func TestMoveGroup(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-group/src"),
	)
	s := testService()
	r, e := s.MoveSymbol(
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
	assertFormatted(t, source)
	assert.StringContains(t, "Second = \"bravo\"", source)
	assert.False(t, strings.Contains(source, "First"))
	moved := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/constant/constant.go",
	)
	assertFormatted(t, moved)
	assert.StringContains(t, "First = \"alfa\"", moved)
	run := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	assertFormatted(t, run)
	assert.StringContains(t, "constant.First + Second", run)
}

func TestMoveFunction(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-function/src"),
	)
	s := testService()
	r, e := s.MoveSymbol(
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
	assertFormatted(t, moved)
	assert.StringContains(t, "package check", moved)
	assert.StringContains(t, "func IsValid(", moved)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go")
	assertFormatted(t, caller)
	assert.StringContains(t, "check.IsValid(\"alfa\")", caller)
	assert.StringContains(t, "example/pkg/check", caller)
}
