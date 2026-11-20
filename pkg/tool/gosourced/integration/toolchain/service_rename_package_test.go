package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRenamePackage(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("package-rename/src"))
	r, e := o.Packages.RenamePackage(
		t.Context(),
		o.Directory,
		"example/pkg/outer/store",
		"depot",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(
		t,
		"pkg/outer/depot/store.go: package store → depot",
		rewrite_tester.JoinedConcerns(r),
	)
	_, e = os.Stat(filepath.Join(o.Directory, "pkg/outer/store"))
	assert.True(t, os.IsNotExist(e))
	moved := o.Read("pkg/outer/depot/store.go")
	rewrite_tester.New(t).Formatted(moved)
	assert.StringContains(t, "package depot", moved)
	run := o.Read("pkg/outer/depot/run.go")
	rewrite_tester.New(t).Formatted(run)
	assert.StringContains(t, "package depot", run)
	assert.StringContains(t, "example/pkg/outer/depot/sub", run)
	tests := o.Read("pkg/outer/depot/store_test.go")
	assert.StringContains(t, "package depot", tests)
	caller := o.Read("pkg/caller/run.go")
	rewrite_tester.New(t).Formatted(caller)
	assert.StringContains(t, "\"example/pkg/outer/depot\"", caller)
	assert.StringContains(t, "v := depot.Store{}", caller)
	assert.False(t, strings.Contains(caller, "store"))
	aliased := o.Read("pkg/aliased/run.go")
	rewrite_tester.New(t).Formatted(aliased)
	assert.StringContains(t, "st \"example/pkg/outer/depot\"", aliased)
	assert.StringContains(t, "return &st.Store{}", aliased)
}

func TestRenamePackageShadowed(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-collision/src"),
	)
	r, e := o.Packages.RenamePackage(
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

func TestRenamePackageSameName(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("package-rename/src"))
	r, e := o.Packages.RenamePackage(
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
