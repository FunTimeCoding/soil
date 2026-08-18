package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestRenameFunction(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("rename-function/src"))
	s := testService()
	r, e := s.Rename(
		d,
		"example/pkg/target",
		"IsGeneratedHeader",
		"IsGenerated",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, len(r.Entries) >= 3)
	declaration := readFixtureFile(t, d, "pkg/target/is_generated_header.go")
	assert.StringContains(t, "func IsGenerated(", declaration)
	inPackage := readFixtureFile(t, d, "pkg/target/check.go")
	assert.StringContains(t, "IsGenerated(content)", inPackage)
	crossPackage := readFixtureFile(t, d, "pkg/caller/run.go")
	assert.StringContains(t, "target.IsGenerated(content)", crossPackage)
}

func TestRenameFunctionSamePackage(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("unexport-function/src"),
	)
	s := testService()
	r, e := s.Rename(
		d,
		"example/pkg/target",
		"IsGenerated",
		"WasGenerated",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	declaration := readFixtureFile(t, d, "pkg/target/is_generated.go")
	assert.StringContains(t, "func WasGenerated(", declaration)
	run := readFixtureFile(t, d, "pkg/target/run.go")
	assert.StringContains(t, "WasGenerated(name)", run)
}

func TestRenameMethod(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("rename-method/src"))
	s := testService()
	r, e := s.Rename(
		d,
		"example/pkg/target",
		"FindByName",
		"LookupByName",
		"Store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, len(r.Entries) >= 3)
	method := readFixtureFile(t, d, "pkg/target/find_by_name.go")
	assert.StringContains(t, "func (s *Store) LookupByName(", method)
	inPackage := readFixtureFile(t, d, "pkg/target/run.go")
	assert.StringContains(t, "v.LookupByName(\"test\")", inPackage)
	crossPackage := readFixtureFile(t, d, "pkg/caller/run.go")
	assert.StringContains(t, "s.LookupByName(\"test\")", crossPackage)
}

func TestRenameToUnexportedBlockedByCrossPackage(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("rename-unexport/src"))
	s := testService()
	r, e := s.Rename(d, "example/pkg/target", "Validate", "check", "", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "would lose access")
	source := readFixtureFile(t, d, "pkg/target/validate.go")
	assert.StringContains(t, "func Validate(", source)
}

func TestRenameToUnexportedAllowedWithinPackage(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("unexport-function/src"),
	)
	s := testService()
	r, e := s.Rename(
		d,
		"example/pkg/target",
		"IsGenerated",
		"wasGenerated",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	declaration := readFixtureFile(t, d, "pkg/target/is_generated.go")
	assert.StringContains(t, "func wasGenerated(", declaration)
}

func TestRenameCollision(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("rename-function/src"))
	s := testService()
	r, e := s.Rename(
		d,
		"example/pkg/target",
		"IsGeneratedHeader",
		"Check",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestRenameSymbolNotFound(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("rename-function/src"))
	s := testService()
	r, e := s.Rename(d, "example/pkg/target", "Missing", "Something", "", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestRenamePackageNotFound(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("rename-function/src"))
	s := testService()
	r, e := s.Rename(
		d,
		"example/pkg/missing",
		"IsGeneratedHeader",
		"IsGenerated",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestRenameReceiverNotFound(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("rename-method/src"))
	s := testService()
	r, e := s.Rename(
		d,
		"example/pkg/target",
		"FindByName",
		"LookupByName",
		"Missing",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestRenameMethodNotFound(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("rename-method/src"))
	s := testService()
	r, e := s.Rename(
		d,
		"example/pkg/target",
		"Missing",
		"Something",
		"Store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestRenameMethodCollision(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("method-collision/src"))
	s := testService()
	r, e := s.Rename(d, "example/pkg/target", "Save", "Load", "Store", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestRenameSameName(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("rename-function/src"))
	s := testService()
	r, e := s.Rename(
		d,
		"example/pkg/target",
		"IsGeneratedHeader",
		"IsGeneratedHeader",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestRenameLongerKeepsComments(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("rename-longer/src"))
	s := testService()
	r, e := s.Rename(
		d,
		"example/pkg/target",
		"Legend",
		"TargetLegendExtended",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := readFixtureFile(t, d, "pkg/target/constant.go")
	assertFormatted(t, source)
	assert.StringContains(t, "TargetLegendExtended = \"--legend\"", source)
	assert.StringContains(t, "// alpha trailing", source)
	assert.StringContains(t, "// bold trailing", source)
	assert.StringContains(t, "// wide trailing", source)
	run := readFixtureFile(t, d, "pkg/target/run.go")
	assertFormatted(t, run)
	assert.StringContains(t, "TargetLegendExtended", run)
}
