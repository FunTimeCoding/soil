package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"strings"
	"testing"
)

func TestRenameModule(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("module-rename/src"))
	r, e := o.Module.RenameModule(
		t.Context(),
		o.Directory,
		"example/lib/v2",
		"example/lib/v3",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	caller := o.Read("pkg/caller/run.go")
	rewrite_tester.New(t).Formatted(caller)
	assert.StringContains(t, "\"example/lib/v3\"", caller)
	assert.False(t, strings.Contains(caller, "example/lib/v2"))
	assert.StringContains(t, "lib.Make(\"alfa\")", caller)
}

func TestRenameModuleBreakage(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("module-break/src"))
	r, e := o.Module.RenameModule(
		t.Context(),
		o.Directory,
		"example/lib/v2",
		"example/lib/v3",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, "Gone no longer exists")
	testutil.AssertBlockedContains(t, r, "size int")
	testutil.AssertBlockedContains(t, r, "→")
	caller := o.Read("pkg/caller/run.go")
	assert.StringContains(t, "\"example/lib/v2\"", caller)
}

func TestRenameModuleTreatsAnyAsEmptyInterface(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("module-any/src"))
	r, e := o.Module.RenameModule(
		t.Context(),
		o.Directory,
		"example/lib/v2",
		"example/lib/v3",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(t, "\"example/lib/v3\"", o.Read("pkg/caller/run.go"))
}

func TestRenameModuleIgnoresParameterAndResultNames(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("module-result-names/src"),
	)
	r, e := o.Module.RenameModule(
		t.Context(),
		o.Directory,
		"example/lib/v2",
		"example/lib/v3",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(t, "\"example/lib/v3\"", o.Read("pkg/caller/run.go"))
}

func TestRenameModuleReadsInternalTypeThroughPublicAlias(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("module-internal-alias/src"),
	)
	r, e := o.Module.RenameModule(
		t.Context(),
		o.Directory,
		"example/lib/v2",
		"example/lib/v3",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	caller := o.Read("pkg/caller/run.go")
	assert.StringContains(t, "\"example/lib/v3\"", caller)
	rewrite_tester.New(t).Compiles(o.Directory)
}

func TestRenameModuleReportsInternalTypeWithoutAlias(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("module-internal-hidden/src"),
	)
	r, e := o.Module.RenameModule(
		t.Context(),
		o.Directory,
		"example/lib/v2",
		"example/lib/v3",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, "internal/scan.Token")
	caller := o.Read("pkg/caller/run.go")
	assert.StringContains(t, "\"example/lib/v2\"", caller)
}

func TestRenameModuleForce(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("module-break/src"))
	r, e := o.Module.RenameModule(
		t.Context(),
		o.Directory,
		"example/lib/v2",
		"example/lib/v3",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	caller := o.Read("pkg/caller/run.go")
	assert.StringContains(t, "\"example/lib/v3\"", caller)
}

func TestRenameModuleDryRun(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("module-rename/src"))
	r, e := o.Module.RenameModule(
		t.Context(),
		o.Directory,
		"example/lib/v2",
		"example/lib/v3",
		false,
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	caller := o.Read("pkg/caller/run.go")
	assert.StringContains(t, "\"example/lib/v2\"", caller)
}

func TestRenameModuleUnused(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("module-rename/src"))
	r, e := o.Module.RenameModule(
		t.Context(),
		o.Directory,
		"example/other/v2",
		"example/other/v3",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, "no references")
}

func TestRenameModuleMissingTarget(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("module-rename/src"))
	r, e := o.Module.RenameModule(
		t.Context(),
		o.Directory,
		"example/lib/v2",
		"example/lib/v4",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, "package not found")
	caller := o.Read("pkg/caller/run.go")
	assert.StringContains(t, "\"example/lib/v2\"", caller)
}
