package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestRenameFunction(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-function/src"),
	)
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
	declaration := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/is_generated_header.go",
	)
	assert.StringContains(t, "func IsGenerated(", declaration)
	inPackage := service_tester.ReadFixtureFile(t, d, "pkg/target/check.go")
	assert.StringContains(t, "IsGenerated(content)", inPackage)
	crossPackage := service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go")
	assert.StringContains(t, "target.IsGenerated(content)", crossPackage)
}

func TestRenameFunctionSamePackage(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("unexport-function/src"),
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
	declaration := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/is_generated.go",
	)
	assert.StringContains(t, "func WasGenerated(", declaration)
	run := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	assert.StringContains(t, "WasGenerated(name)", run)
}

func TestRenameMethod(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-method/src"),
	)
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
	method := service_tester.ReadFixtureFile(t, d, "pkg/target/find_by_name.go")
	assert.StringContains(t, "func (s *Store) LookupByName(", method)
	inPackage := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	assert.StringContains(t, "v.LookupByName(\"test\")", inPackage)
	crossPackage := service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go")
	assert.StringContains(t, "s.LookupByName(\"test\")", crossPackage)
}

func TestRenameToUnexportedBlockedByCrossPackage(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-unexport/src"),
	)
	s := testService()
	r, e := s.Rename(d, "example/pkg/target", "Validate", "check", "", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "would lose access")
	source := service_tester.ReadFixtureFile(t, d, "pkg/target/validate.go")
	assert.StringContains(t, "func Validate(", source)
}

func TestRenameToUnexportedAllowedWithinPackage(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("unexport-function/src"),
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
	declaration := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/is_generated.go",
	)
	assert.StringContains(t, "func wasGenerated(", declaration)
}

func TestRenameCollision(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-function/src"),
	)
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
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-function/src"),
	)
	s := testService()
	r, e := s.Rename(d, "example/pkg/target", "Missing", "Something", "", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestRenamePackageNotFound(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-function/src"),
	)
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
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-method/src"),
	)
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
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-method/src"),
	)
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
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("method-collision/src"),
	)
	s := testService()
	r, e := s.Rename(d, "example/pkg/target", "Save", "Load", "Store", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestRenameSameName(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-function/src"),
	)
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
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-longer/src"),
	)
	s := testService()
	r, e := s.Rename(
		d,
		"example/pkg/target/constant",
		"Legend",
		"TargetLegendExtended",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/constant/constant.go",
	)
	assertFormatted(t, source)
	assert.StringContains(t, "TargetLegendExtended = \"--legend\"", source)
	assert.StringContains(t, "// alpha trailing", source)
	assert.StringContains(t, "// bold trailing", source)
	assert.StringContains(t, "// wide trailing", source)
	run := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	assertFormatted(t, run)
	assert.StringContains(t, "TargetLegendExtended", run)
}
