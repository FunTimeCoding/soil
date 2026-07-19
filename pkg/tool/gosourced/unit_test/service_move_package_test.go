package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMovePackage(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("package-move/src"))
	s := testService()
	r, e := s.MovePackage(
		d,
		"example/pkg/outer/store",
		"example/pkg/store",
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	_, e = os.Stat(filepath.Join(d, "pkg/outer/store"))
	assert.True(t, os.IsNotExist(e))
	moved := readFixtureFile(t, d, "pkg/store/store.go")
	assert.StringContains(t, "package store", moved)
	run := readFixtureFile(t, d, "pkg/store/run.go")
	assertFormatted(t, run)
	assert.StringContains(t, "example/pkg/store/sub", run)
	assert.False(t, strings.Contains(run, "outer"))
	caller := readFixtureFile(t, d, "pkg/caller/run.go")
	assertFormatted(t, caller)
	assert.StringContains(t, "\"example/pkg/store\"", caller)
	assert.StringContains(t, "\"example/pkg/store/sub\"", caller)
	assert.False(t, strings.Contains(caller, "outer"))
}

func TestMovePackageBaseMismatch(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("package-move/src"))
	s := testService()
	r, e := s.MovePackage(
		d,
		"example/pkg/outer/store",
		"example/pkg/depot",
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "name would change")
}

func TestMovePackageTargetExists(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("package-move/src"))
	e := os.MkdirAll(filepath.Join(d, "pkg/store"), 0755)
	assert.FatalOnError(t, e)
	s := testService()
	r, e := s.MovePackage(
		d,
		"example/pkg/outer/store",
		"example/pkg/store",
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestMovePackageIntoItself(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("package-move/src"))
	s := testService()
	r, e := s.MovePackage(
		d,
		"example/pkg/outer/store",
		"example/pkg/outer/store/inner/store",
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "into itself")
}

func TestMovePackageNotFound(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("package-move/src"))
	s := testService()
	r, e := s.MovePackage(
		d,
		"example/pkg/outer/missing",
		"example/pkg/missing",
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
}
