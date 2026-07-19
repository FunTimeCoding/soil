package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBatchFile(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("batch-file/src"))
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/target",
		nil,
		"pkg/target/constant.go",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	_, e = os.Stat(filepath.Join(d, "pkg/target/constant.go"))
	assert.True(t, os.IsNotExist(e))
	moved := readFixtureFile(t, d, "pkg/target/constant/constant.go")
	assertFormatted(t, moved)
	assert.Integer(t, 1, strings.Count(moved, "const ("))
	assert.StringContains(t, "First  = \"alfa\"", moved)
	assert.StringContains(t, "Second = \"bravo\"", moved)
	assert.StringContains(t, "Third  = \"charlie\"", moved)
	run := readFixtureFile(t, d, "pkg/target/run.go")
	assertFormatted(t, run)
	assert.StringContains(
		t,
		"constant.First + constant.Second + constant.Third",
		run,
	)
}

func TestBatchEnum(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("batch-enum/src"))
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/target",
		nil,
		"pkg/target/constant.go",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := readFixtureFile(t, d, "pkg/target/constant/constant.go")
	assertFormatted(t, moved)
	assert.StringContains(t, "type Mode int", moved)
	assert.StringContains(t, "Off Mode = iota", moved)
	assert.False(t, strings.Contains(moved, "constant.Mode"))
	run := readFixtureFile(t, d, "pkg/target/run.go")
	assertFormatted(t, run)
	assert.StringContains(t, "func Run() constant.Mode", run)
	assert.StringContains(t, "return constant.On", run)
	caller := readFixtureFile(t, d, "pkg/caller/run.go")
	assertFormatted(t, caller)
	assert.StringContains(t, "func Run() constant.Mode", caller)
	assert.StringContains(t, "constant.Off == 0", caller)
	assert.False(t, strings.Contains(caller, "target."))
}

func TestBatchSubset(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("batch-file/src"))
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/target",
		[]string{"first", "second"},
		"",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := readFixtureFile(t, d, "pkg/target/constant.go")
	assertFormatted(t, source)
	assert.StringContains(t, "third = \"charlie\"", source)
	assert.False(t, strings.Contains(source, "first"))
	moved := readFixtureFile(t, d, "pkg/target/constant/constant.go")
	assertFormatted(t, moved)
	assert.StringContains(t, "First  = \"alfa\"", moved)
	assert.StringContains(t, "Second = \"bravo\"", moved)
	run := readFixtureFile(t, d, "pkg/target/run.go")
	assert.StringContains(
		t,
		"constant.First + constant.Second + third",
		run,
	)
}

func TestBatchCollision(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("batch-collision/src"))
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/target",
		nil,
		"pkg/target/constant.go",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "both move as Mode")
}

func TestBatchDependencyTogether(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("move-dependency/src"))
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/target",
		[]string{"prefix", "composed"},
		"",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := readFixtureFile(t, d, "pkg/target/constant/constant.go")
	assertFormatted(t, moved)
	assert.StringContains(t, "Composed = Prefix + \"bravo\"", moved)
}

func TestBatchMultiName(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("batch-multi-name/src"))
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/target",
		nil,
		"pkg/target/constant.go",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := readFixtureFile(t, d, "pkg/target/constant/constant.go")
	assertFormatted(t, moved)
	assert.StringContains(t, "Alfa, Bravo = \"a\", \"b\"", moved)
	run := readFixtureFile(t, d, "pkg/target/run.go")
	assert.StringContains(t, "constant.Alfa + constant.Bravo", run)
}

func TestBatchMultiNamePartial(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("batch-multi-name/src"))
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/target",
		[]string{"alfa"},
		"",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "move them together")
}

func TestBatchMethodRefused(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("batch-method/src"))
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/target",
		nil,
		"pkg/target/store.go",
		"example/pkg/target/constant",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "methods cannot move")
}

func TestBackReferenceQualified(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("back-reference/src"))
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/source",
		[]string{"Mover"},
		"",
		"example/pkg/target",
		"mover.go",
		true,
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := readFixtureFile(t, d, "pkg/target/mover.go")
	assertFormatted(t, moved)
	assert.StringContains(t, "example/pkg/source", moved)
	assert.StringContains(t, "source.Stayer()", moved)
	source := readFixtureFile(t, d, "pkg/source/source.go")
	assertFormatted(t, source)
	assert.StringContains(t, "func Stayer()", source)
	assert.False(t, strings.Contains(source, "func Mover"))
}

func TestBackReferenceRefusedWithoutFlag(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("back-reference/src"))
	s := testService()
	r, e := s.MoveSymbols(
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
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "package-local")
}

func TestBackReferenceCycleRefused(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("cycle/src"))
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/source",
		[]string{"Mover"},
		"",
		"example/pkg/target",
		"mover.go",
		true,
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "cycle")
}

func TestBackReferenceUnexportedRefused(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("back-reference-unexported/src"),
	)
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/source",
		[]string{"Mover"},
		"",
		"example/pkg/target",
		"mover.go",
		true,
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "unexported")
}
