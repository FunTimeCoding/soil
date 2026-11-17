package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestRenameOntoImportedNameRefuses(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("rename-scope/src"))
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"Render",
		"bravo",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "bravo is taken at")
	testutil.AssertBlockedContains(t, r, "pkg/alfa/show.go:3")
	rewrite_tester.Compiles(t, o.Directory)
}

func TestRenameOntoLocalAtCallSiteRefuses(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("rename-scope/src"))
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"Count",
		"size",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "size is taken at")
	testutil.AssertBlockedContains(t, r, "pkg/alfa/total.go:6")
	rewrite_tester.Compiles(t, o.Directory)
}

func TestRenameOntoUsedBuiltinRefuses(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("rename-scope/src"))
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"Width",
		"len",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "len is a built-in name")
	rewrite_tester.Compiles(t, o.Directory)
}

func TestRenameBesideUnrelatedLocalBuilds(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("rename-scope/src"))
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"Count",
		"amount",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	rewrite_tester.Compiles(t, o.Directory)
}
