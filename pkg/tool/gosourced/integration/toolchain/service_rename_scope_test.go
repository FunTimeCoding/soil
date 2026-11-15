package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestRenameOntoImportedNameRefuses(t *testing.T) {
	d := service_tester.Prepare(t, "rename-scope/src")
	r, e := service_tester.NewService(t).Symbol.Rename(
		t.Context(),
		d,
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
	rewrite_tester.Compiles(t, d)
}

func TestRenameOntoLocalAtCallSiteRefuses(t *testing.T) {
	d := service_tester.Prepare(t, "rename-scope/src")
	r, e := service_tester.NewService(t).Symbol.Rename(
		t.Context(),
		d,
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
	rewrite_tester.Compiles(t, d)
}

func TestRenameOntoUsedBuiltinRefuses(t *testing.T) {
	d := service_tester.Prepare(t, "rename-scope/src")
	r, e := service_tester.NewService(t).Symbol.Rename(
		t.Context(),
		d,
		"example/pkg/alfa",
		"Width",
		"len",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "len is a built-in name")
	rewrite_tester.Compiles(t, d)
}

func TestRenameBesideUnrelatedLocalBuilds(t *testing.T) {
	d := service_tester.Prepare(t, "rename-scope/src")
	r, e := service_tester.NewService(t).Symbol.Rename(
		t.Context(),
		d,
		"example/pkg/alfa",
		"Count",
		"amount",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	rewrite_tester.Compiles(t, d)
}
