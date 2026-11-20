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

func TestMovePackage(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("package-move/src"))
	r, e := o.Packages.MovePackage(
		t.Context(),
		o.Directory,
		"example/pkg/outer/store",
		"example/pkg/store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	_, e = os.Stat(filepath.Join(o.Directory, "pkg/outer/store"))
	assert.True(t, os.IsNotExist(e))
	moved := o.Read("pkg/store/store.go")
	assert.StringContains(t, "package store", moved)
	run := o.Read("pkg/store/run.go")
	rewrite_tester.New(t).Formatted(run)
	assert.StringContains(t, "example/pkg/store/sub", run)
	assert.False(t, strings.Contains(run, "outer"))
	caller := o.Read("pkg/caller/run.go")
	rewrite_tester.New(t).Formatted(caller)
	assert.StringContains(t, "\"example/pkg/store\"", caller)
	assert.StringContains(t, "\"example/pkg/store/sub\"", caller)
	assert.False(t, strings.Contains(caller, "outer"))
}

func TestMovePackageTargetExists(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("package-move/src"))
	e := os.MkdirAll(filepath.Join(o.Directory, "pkg/store"), 0755)
	assert.FatalOnError(t, e)
	r, e := o.Packages.MovePackage(
		t.Context(),
		o.Directory,
		"example/pkg/outer/store",
		"example/pkg/store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestMovePackageNotFound(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("package-move/src"))
	r, e := o.Packages.MovePackage(
		t.Context(),
		o.Directory,
		"example/pkg/outer/missing",
		"example/pkg/missing",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}
