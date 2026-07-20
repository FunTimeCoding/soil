package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"os"
	"path/filepath"
	"testing"
)

func testService() *service.Service {
	return service.New(inventory.New())
}

func readFixtureFile(
	t *testing.T,
	d string,
	path string,
) string {
	t.Helper()
	b, e := os.ReadFile(filepath.Join(d, path))
	assert.FatalOnError(t, e)

	return string(b)
}

func TestUnexportFunction(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("unexport-function/src"),
	)
	s := testService()
	r, e := s.ChangeVisibility(d, "IsGenerated", "example/pkg/target", "")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, len(r.Entries) >= 2)
	helper := readFixtureFile(t, d, "pkg/target/is_generated.go")
	assert.StringContains(t, "func isGenerated(", helper)
	run := readFixtureFile(t, d, "pkg/target/run.go")
	assert.StringContains(t, "isGenerated(name)", run)
}

func TestExportFunction(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("export-function/src"),
	)
	s := testService()
	r, e := s.ChangeVisibility(d, "isGenerated", "example/pkg/target", "")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, len(r.Entries) >= 2)
	helper := readFixtureFile(t, d, "pkg/target/is_generated.go")
	assert.StringContains(t, "func IsGenerated(", helper)
	run := readFixtureFile(t, d, "pkg/target/run.go")
	assert.StringContains(t, "IsGenerated(\"test\")", run)
}

func TestUnexportMethod(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("unexport-method/src"),
	)
	s := testService()
	r, e := s.ChangeVisibility(d, "Save", "example/pkg/target", "Store")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, len(r.Entries) >= 2)
	save := readFixtureFile(t, d, "pkg/target/save.go")
	assert.StringContains(t, "func (s *Store) save(", save)
	run := readFixtureFile(t, d, "pkg/target/run.go")
	assert.StringContains(t, "v.save(\"test\")", run)
}

func TestCollisionDetection(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("collision/src"))
	s := testService()
	r, e := s.ChangeVisibility(d, "IsGenerated", "example/pkg/target", "")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestSymbolNotFound(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("unexport-function/src"),
	)
	s := testService()
	r, e := s.ChangeVisibility(d, "Missing", "example/pkg/target", "")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestPackageNotFound(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("unexport-function/src"),
	)
	s := testService()
	r, e := s.ChangeVisibility(d, "IsGenerated", "example/pkg/missing", "")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}

func TestUnexportBlockedByCrossPackageCaller(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("cross-package/src"))
	s := testService()
	r, e := s.ChangeVisibility(d, "IsValid", "example/pkg/target", "")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "would lose access")
	helper := readFixtureFile(t, d, "pkg/target/is_valid.go")
	assert.StringContains(t, "func IsValid(", helper)
}
