package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func assertFormatted(
	t *testing.T,
	content string,
) {
	t.Helper()
	formatted, e := format.Source([]byte(content))
	assert.FatalOnError(t, e)
	assert.String(t, string(formatted), content)
}

func TestMoveConstant(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("move-constant/src"))
	s := testService()
	r, e := s.MoveSymbol(
		d,
		"example/pkg/target",
		"itemFields",
		"example/pkg/target/constant",
		"constant.go",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	_, e = os.Stat(filepath.Join(d, "pkg/target/constant.go"))
	assert.True(t, os.IsNotExist(e))
	moved := readFixtureFile(t, d, "pkg/target/constant/constant.go")
	assertFormatted(t, moved)
	assert.StringContains(t, "ItemFields = \"alfa\"", moved)
	assert.StringContains(t, "Host = \"bravo\"", moved)
	run := readFixtureFile(t, d, "pkg/target/run.go")
	assertFormatted(t, run)
	assert.StringContains(t, "return constant.ItemFields", run)
	assert.StringContains(t, "example/pkg/target/constant", run)
}

func TestMoveCrossPackage(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("move-cross-package/src"),
	)
	s := testService()
	r, e := s.MoveSymbol(
		d,
		"example/pkg/target",
		"Mode",
		"example/pkg/target/constant",
		"constant.go",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	run := readFixtureFile(t, d, "pkg/target/run.go")
	assertFormatted(t, run)
	assert.StringContains(t, "return constant.Mode", run)
	caller := readFixtureFile(t, d, "pkg/caller/run.go")
	assertFormatted(t, caller)
	assert.StringContains(t, "return constant.Mode", caller)
	assert.StringContains(t, "example/pkg/target/constant", caller)
	assert.False(t, strings.Contains(caller, "target.Mode"))
}

func TestMoveAlias(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("move-alias/src"))
	s := testService()
	r, e := s.MoveSymbol(
		d,
		"example/pkg/target",
		"Mode",
		"example/pkg/target/constant",
		"constant.go",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	caller := readFixtureFile(t, d, "pkg/caller/run.go")
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
	d := testutil.PrepareTestPackage(t, serviceTestdata("move-create/src"))
	s := testService()
	r, e := s.MoveSymbol(
		d,
		"example/pkg/target",
		"itemFields",
		"example/pkg/target/constant",
		"constant.go",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestMoveCreate(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("move-create/src"))
	s := testService()
	r, e := s.MoveSymbol(
		d,
		"example/pkg/target",
		"itemFields",
		"example/pkg/target/constant",
		"constant.go",
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := readFixtureFile(t, d, "pkg/target/constant/constant.go")
	assertFormatted(t, moved)
	assert.StringContains(t, "package constant", moved)
	assert.StringContains(t, "ItemFields = \"alfa\"", moved)
	run := readFixtureFile(t, d, "pkg/target/run.go")
	assertFormatted(t, run)
	assert.StringContains(t, "return constant.ItemFields", run)
}

func TestMoveCollision(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("move-collision/src"))
	s := testService()
	r, e := s.MoveSymbol(
		d,
		"example/pkg/target",
		"itemFields",
		"example/pkg/target/constant",
		"constant.go",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestMoveCycle(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("move-cycle/src"))
	s := testService()
	r, e := s.MoveSymbol(
		d,
		"example/pkg/target",
		"Mode",
		"example/pkg/other",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "cycle")
}

func TestMoveDependency(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("move-dependency/src"),
	)
	s := testService()
	r, e := s.MoveSymbol(
		d,
		"example/pkg/target",
		"composed",
		"example/pkg/target/constant",
		"constant.go",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "package-local")
	testutil.AssertBlockedContains(t, r, "prefix")
}

func TestMoveGroup(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("move-group/src"))
	s := testService()
	r, e := s.MoveSymbol(
		d,
		"example/pkg/target",
		"First",
		"example/pkg/target/constant",
		"constant.go",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := readFixtureFile(t, d, "pkg/target/constant.go")
	assertFormatted(t, source)
	assert.StringContains(t, "Second = \"bravo\"", source)
	assert.False(t, strings.Contains(source, "First"))
	moved := readFixtureFile(t, d, "pkg/target/constant/constant.go")
	assertFormatted(t, moved)
	assert.StringContains(t, "First = \"alfa\"", moved)
	run := readFixtureFile(t, d, "pkg/target/run.go")
	assertFormatted(t, run)
	assert.StringContains(t, "constant.First + Second", run)
}

func TestMoveFunction(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("move-function/src"))
	s := testService()
	r, e := s.MoveSymbol(
		d,
		"example/pkg/target",
		"IsValid",
		"example/pkg/check",
		"",
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	_, e = os.Stat(filepath.Join(d, "pkg/target/is_valid.go"))
	assert.True(t, os.IsNotExist(e))
	moved := readFixtureFile(t, d, "pkg/check/is_valid.go")
	assertFormatted(t, moved)
	assert.StringContains(t, "package check", moved)
	assert.StringContains(t, "func IsValid(", moved)
	caller := readFixtureFile(t, d, "pkg/caller/run.go")
	assertFormatted(t, caller)
	assert.StringContains(t, "check.IsValid(\"alfa\")", caller)
	assert.StringContains(t, "example/pkg/check", caller)
}
