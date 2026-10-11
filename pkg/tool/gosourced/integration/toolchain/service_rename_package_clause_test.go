package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestRenamePackageClause(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("package-rename/src"),
	)
	s := service_tester.NewService()
	r, e := s.RenamePackageClause(d, "example/pkg/outer/store", "depot", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	reported := joinedConcerns(r)
	assert.StringContains(
		t,
		"pkg/outer/store/store.go: package store → depot",
		reported,
	)
	assert.StringContains(
		t,
		"pkg/outer/store/store_test.go: package store → depot",
		reported,
	)
	_, e = os.Stat(filepath.Join(d, "pkg/outer/store"))
	assert.FatalOnError(t, e)
	clause := service_tester.ReadFixtureFile(t, d, "pkg/outer/store/store.go")
	assertFormatted(t, clause)
	assert.StringContains(t, "package depot", clause)
	run := service_tester.ReadFixtureFile(t, d, "pkg/outer/store/run.go")
	assertFormatted(t, run)
	assert.StringContains(t, "package depot", run)
	assert.StringContains(t, "example/pkg/outer/store/sub", run)
	tests := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/outer/store/store_test.go",
	)
	assert.StringContains(t, "package depot", tests)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go")
	assertFormatted(t, caller)
	assert.StringContains(t, "\"example/pkg/outer/store\"", caller)
	assert.StringContains(t, "v := depot.Store{}", caller)
	aliased := service_tester.ReadFixtureFile(t, d, "pkg/aliased/run.go")
	assertFormatted(t, aliased)
	assert.StringContains(t, "st \"example/pkg/outer/store\"", aliased)
	assert.StringContains(t, "return &st.Store{}", aliased)
}

func TestRenamePackageClauseShadowed(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-collision/src"),
	)
	s := service_tester.NewService()
	r, e := s.RenamePackageClause(d, "example/pkg/outer/store", "depot", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "taken")
}

func TestRenamePackageClauseSameName(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("package-rename/src"),
	)
	s := service_tester.NewService()
	r, e := s.RenamePackageClause(d, "example/pkg/outer/store", "store", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already named")
}

func TestRenamePackageClauseReportsATestOnlyPackage(t *testing.T) {
	d := t.TempDir()
	testutil.WriteModFile(t, d, "example")
	testutil.WriteFile(
		t,
		d,
		"pkg/alfa/alfa_test.go",
		"package guard\n\nimport \"testing\"\n\nfunc TestAlfa(t *testing.T) {}\n",
	)
	s := service_tester.NewService()
	r, e := s.RenamePackageClause(d, "example/pkg/alfa", "unit", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"pkg/alfa/alfa_test.go: package guard → unit",
		joinedConcerns(r),
	)
	assert.StringContains(
		t,
		"package unit",
		service_tester.ReadFixtureFile(t, d, "pkg/alfa/alfa_test.go"),
	)
}
