package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMovePackage(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("package-move/src"),
	)
	s := testService()
	r, e := s.MovePackage(
		d,
		"example/pkg/outer/store",
		"example/pkg/store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	_, e = os.Stat(filepath.Join(d, "pkg/outer/store"))
	assert.True(t, os.IsNotExist(e))
	moved := service_tester.ReadFixtureFile(t, d, "pkg/store/store.go")
	assert.StringContains(t, "package store", moved)
	run := service_tester.ReadFixtureFile(t, d, "pkg/store/run.go")
	assertFormatted(t, run)
	assert.StringContains(t, "example/pkg/store/sub", run)
	assert.False(t, strings.Contains(run, "outer"))
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go")
	assertFormatted(t, caller)
	assert.StringContains(t, "\"example/pkg/store\"", caller)
	assert.StringContains(t, "\"example/pkg/store/sub\"", caller)
	assert.False(t, strings.Contains(caller, "outer"))
}

func TestMovePackageBaseMismatch(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("package-move/src"),
	)
	s := testService()
	r, e := s.MovePackage(
		d,
		"example/pkg/outer/store",
		"example/pkg/depot",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "name would change")
}

func TestMovePackageTargetExists(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("package-move/src"),
	)
	e := os.MkdirAll(filepath.Join(d, "pkg/store"), 0755)
	assert.FatalOnError(t, e)
	s := testService()
	r, e := s.MovePackage(
		d,
		"example/pkg/outer/store",
		"example/pkg/store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestMovePackageIntoItself(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("package-move/src"),
	)
	s := testService()
	r, e := s.MovePackage(
		d,
		"example/pkg/outer/store",
		"example/pkg/outer/store/inner/store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "into itself")
}

func TestMovePackageNotFound(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("package-move/src"),
	)
	s := testService()
	r, e := s.MovePackage(
		d,
		"example/pkg/outer/missing",
		"example/pkg/missing",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}
