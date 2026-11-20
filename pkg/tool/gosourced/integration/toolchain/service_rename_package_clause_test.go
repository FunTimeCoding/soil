package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestRenamePackageClause(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("package-rename/src"))
	r, e := o.Packages.RenamePackageClause(
		t.Context(),
		o.Directory,
		"example/pkg/outer/store",
		"depot",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	reported := rewrite_tester.JoinedConcerns(r)
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
	_, e = os.Stat(filepath.Join(o.Directory, "pkg/outer/store"))
	assert.FatalOnError(t, e)
	clause := o.Read("pkg/outer/store/store.go")
	rewrite_tester.New(t).Formatted(clause)
	assert.StringContains(t, "package depot", clause)
	run := o.Read("pkg/outer/store/run.go")
	rewrite_tester.New(t).Formatted(run)
	assert.StringContains(t, "package depot", run)
	assert.StringContains(t, "example/pkg/outer/store/sub", run)
	tests := o.Read("pkg/outer/store/store_test.go")
	assert.StringContains(t, "package depot", tests)
	caller := o.Read("pkg/caller/run.go")
	rewrite_tester.New(t).Formatted(caller)
	assert.StringContains(t, "\"example/pkg/outer/store\"", caller)
	assert.StringContains(t, "v := depot.Store{}", caller)
	aliased := o.Read("pkg/aliased/run.go")
	rewrite_tester.New(t).Formatted(aliased)
	assert.StringContains(t, "st \"example/pkg/outer/store\"", aliased)
	assert.StringContains(t, "return &st.Store{}", aliased)
}

func TestRenamePackageClauseShadowed(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-collision/src"),
	)
	r, e := o.Packages.RenamePackageClause(
		t.Context(),
		o.Directory,
		"example/pkg/outer/store",
		"depot",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "taken")
}

func TestRenamePackageClauseSameName(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("package-rename/src"))
	r, e := o.Packages.RenamePackageClause(
		t.Context(),
		o.Directory,
		"example/pkg/outer/store",
		"store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already named")
}

func TestRenamePackageClauseReportsTestOnlyPackage(t *testing.T) {
	o := service_tester.New(t, service_tester.WithDirectory(t.TempDir()))
	testutil.WriteModFile(t, o.Directory, "example")
	system_tester.New(t).Write(
		o.Directory,
		"pkg/alfa/alfa_test.go",
		"package guard\n\nimport \"testing\"\n\nfunc TestAlfa(t *testing.T) {}\n",
	)
	r, e := o.Packages.RenamePackageClause(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		constant.UnitSegment,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"pkg/alfa/alfa_test.go: package guard → unit",
		rewrite_tester.JoinedConcerns(r),
	)
	assert.StringContains(t, "package unit", o.Read("pkg/alfa/alfa_test.go"))
}
