package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestUnexportFunction(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("unexport-function/src"),
	)
	s := testService()
	r, e := s.ChangeVisibility(
		d,
		"IsGenerated",
		"example/pkg/target",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, len(r.Entries) >= 2)
	helper := service_tester.ReadFixtureFile(t, d, "pkg/target/is_generated.go")
	assert.StringContains(t, "func isGenerated(", helper)
	run := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	assert.StringContains(t, "isGenerated(name)", run)
}

func TestExportFunction(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("export-function/src"),
	)
	s := testService()
	r, e := s.ChangeVisibility(
		d,
		"isGenerated",
		"example/pkg/target",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, len(r.Entries) >= 2)
	helper := service_tester.ReadFixtureFile(t, d, "pkg/target/is_generated.go")
	assert.StringContains(t, "func IsGenerated(", helper)
	run := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	assert.StringContains(t, "IsGenerated(\"test\")", run)
}

func TestUnexportMethod(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("unexport-method/src"),
	)
	s := testService()
	r, e := s.ChangeVisibility(d, "Save", "example/pkg/target", "Store", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, len(r.Entries) >= 2)
	save := service_tester.ReadFixtureFile(t, d, "pkg/target/save.go")
	assert.StringContains(t, "func (s *Store) save(", save)
	run := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	assert.StringContains(t, "v.save(\"test\")", run)
}

func TestCollisionDetection(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("collision/src"),
	)
	s := testService()
	r, e := s.ChangeVisibility(
		d,
		"IsGenerated",
		"example/pkg/target",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestSymbolNotFound(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("unexport-function/src"),
	)
	s := testService()
	r, e := s.ChangeVisibility(d, "Missing", "example/pkg/target", "", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestPackageNotFound(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("unexport-function/src"),
	)
	s := testService()
	r, e := s.ChangeVisibility(
		d,
		"IsGenerated",
		"example/pkg/missing",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestUnexportBlockedByCrossPackageCaller(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("cross-package/src"),
	)
	s := testService()
	r, e := s.ChangeVisibility(d, "IsValid", "example/pkg/target", "", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "would lose access")
	helper := service_tester.ReadFixtureFile(t, d, "pkg/target/is_valid.go")
	assert.StringContains(t, "func IsValid(", helper)
}
