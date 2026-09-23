package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"strings"
	"testing"
)

func TestRenameModule(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("module-rename/src"),
	)
	s := testService()
	r, e := s.RenameModule(d, "example/lib/v2", "example/lib/v3", false, false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go")
	assertFormatted(t, caller)
	assert.StringContains(t, "\"example/lib/v3\"", caller)
	assert.False(t, strings.Contains(caller, "example/lib/v2"))
	assert.StringContains(t, "lib.Make(\"alfa\")", caller)
}

func TestRenameModuleBreakage(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("module-break/src"),
	)
	s := testService()
	r, e := s.RenameModule(d, "example/lib/v2", "example/lib/v3", false, false)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, "Gone no longer exists")
	testutil.AssertBlockedContains(t, r, "size int")
	testutil.AssertBlockedContains(t, r, "→")
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go")
	assert.StringContains(t, "\"example/lib/v2\"", caller)
}

func TestRenameModuleForce(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("module-break/src"),
	)
	s := testService()
	r, e := s.RenameModule(d, "example/lib/v2", "example/lib/v3", true, false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go")
	assert.StringContains(t, "\"example/lib/v3\"", caller)
}

func TestRenameModuleDryRun(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("module-rename/src"),
	)
	s := testService()
	r, e := s.RenameModule(d, "example/lib/v2", "example/lib/v3", false, true)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go")
	assert.StringContains(t, "\"example/lib/v2\"", caller)
}

func TestRenameModuleUnused(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("module-rename/src"),
	)
	s := testService()
	r, e := s.RenameModule(
		d,
		"example/other/v2",
		"example/other/v3",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, "no references")
}

func TestRenameModuleMissingTarget(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("module-rename/src"),
	)
	s := testService()
	r, e := s.RenameModule(d, "example/lib/v2", "example/lib/v4", false, false)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, "package not found")
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go")
	assert.StringContains(t, "\"example/lib/v2\"", caller)
}

func TestRenameModuleSamePath(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("module-rename/src"),
	)
	s := testService()
	r, e := s.RenameModule(d, "example/lib/v2", "example/lib/v2", false, false)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, "same")
}
