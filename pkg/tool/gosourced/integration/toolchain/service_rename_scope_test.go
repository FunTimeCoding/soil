package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestRenameOntoAnImportedNameRefuses(t *testing.T) {
	d := renameScopeFixture(t)
	r, e := service_tester.NewService().Rename(
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
	assertCompiles(t, d)
}

func TestRenameOntoALocalAtACallSiteRefuses(t *testing.T) {
	d := renameScopeFixture(t)
	r, e := service_tester.NewService().Rename(
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
	assertCompiles(t, d)
}

func TestRenameOntoAUsedBuiltinRefuses(t *testing.T) {
	d := renameScopeFixture(t)
	r, e := service_tester.NewService().Rename(
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
	assertCompiles(t, d)
}

func TestRenameBesideAnUnrelatedLocalBuilds(t *testing.T) {
	d := renameScopeFixture(t)
	r, e := service_tester.NewService().Rename(
		d,
		"example/pkg/alfa",
		"Count",
		"amount",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assertCompiles(t, d)
}
