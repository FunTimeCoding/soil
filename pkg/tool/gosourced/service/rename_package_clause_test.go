package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"os"
	"path/filepath"
	"testing"
)

func TestRenamePackageClause(t *testing.T) {
	d := testutil.PrepareTestPackage(t, "testdata/package-rename/src")
	s := testService()
	r, e := s.RenamePackageClause(d, "example/pkg/outer/store", "depot")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	_, e = os.Stat(filepath.Join(d, "pkg/outer/store"))
	assert.FatalOnError(t, e)
	clause := readFixtureFile(t, d, "pkg/outer/store/store.go")
	assertFormatted(t, clause)
	assert.StringContains(t, "package depot", clause)
	run := readFixtureFile(t, d, "pkg/outer/store/run.go")
	assertFormatted(t, run)
	assert.StringContains(t, "package depot", run)
	assert.StringContains(t, "example/pkg/outer/store/sub", run)
	tests := readFixtureFile(t, d, "pkg/outer/store/store_test.go")
	assert.StringContains(t, "package depot", tests)
	caller := readFixtureFile(t, d, "pkg/caller/run.go")
	assertFormatted(t, caller)
	assert.StringContains(t, "\"example/pkg/outer/store\"", caller)
	assert.StringContains(t, "v := depot.Store{}", caller)
	aliased := readFixtureFile(t, d, "pkg/aliased/run.go")
	assertFormatted(t, aliased)
	assert.StringContains(
		t,
		"st \"example/pkg/outer/store\"",
		aliased,
	)
	assert.StringContains(t, "return &st.Store{}", aliased)
}

func TestRenamePackageClauseShadowed(t *testing.T) {
	d := testutil.PrepareTestPackage(t, "testdata/rename-collision/src")
	s := testService()
	r, e := s.RenamePackageClause(d, "example/pkg/outer/store", "depot")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "taken")
}

func TestRenamePackageClauseSameName(t *testing.T) {
	d := testutil.PrepareTestPackage(t, "testdata/package-rename/src")
	s := testService()
	r, e := s.RenamePackageClause(d, "example/pkg/outer/store", "store")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already named")
}
