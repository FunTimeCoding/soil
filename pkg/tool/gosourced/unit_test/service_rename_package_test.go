package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRenamePackage(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("package-rename/src"))
	s := testService()
	r, e := s.RenamePackage(d, "example/pkg/outer/store", "depot")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	_, e = os.Stat(filepath.Join(d, "pkg/outer/store"))
	assert.True(t, os.IsNotExist(e))
	moved := readFixtureFile(t, d, "pkg/outer/depot/store.go")
	assertFormatted(t, moved)
	assert.StringContains(t, "package depot", moved)
	run := readFixtureFile(t, d, "pkg/outer/depot/run.go")
	assertFormatted(t, run)
	assert.StringContains(t, "package depot", run)
	assert.StringContains(t, "example/pkg/outer/depot/sub", run)
	tests := readFixtureFile(t, d, "pkg/outer/depot/store_test.go")
	assert.StringContains(t, "package depot", tests)
	caller := readFixtureFile(t, d, "pkg/caller/run.go")
	assertFormatted(t, caller)
	assert.StringContains(t, "\"example/pkg/outer/depot\"", caller)
	assert.StringContains(t, "v := depot.Store{}", caller)
	assert.False(t, strings.Contains(caller, "store"))
	aliased := readFixtureFile(t, d, "pkg/aliased/run.go")
	assertFormatted(t, aliased)
	assert.StringContains(
		t,
		"st \"example/pkg/outer/depot\"",
		aliased,
	)
	assert.StringContains(t, "return &st.Store{}", aliased)
}

func TestRenamePackageShadowed(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("rename-collision/src"))
	s := testService()
	r, e := s.RenamePackage(d, "example/pkg/outer/store", "depot")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "taken")
}

func TestRenamePackageSameName(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("package-rename/src"))
	s := testService()
	r, e := s.RenamePackage(d, "example/pkg/outer/store", "store")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already named")
}

func TestRenamePackageInvalidName(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("package-rename/src"))
	s := testService()
	r, e := s.RenamePackage(d, "example/pkg/outer/store", "9depot")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not a valid")
}
