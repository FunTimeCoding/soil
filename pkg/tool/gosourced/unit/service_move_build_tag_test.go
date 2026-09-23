package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"strings"
	"testing"
)

func TestMoveBuildTagCarry(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-build-tag/src"),
	)
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/tagged",
		[]string{"Flag"},
		"",
		"example/pkg/tagged/hold",
		"hold.go",
		true,
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := service_tester.ReadFixtureFile(t, d, "pkg/tagged/hold/hold.go")
	assertFormatted(t, moved)
	assert.True(t, strings.HasPrefix(moved, "//go:build local\n\npackage hold"))
	assert.StringContains(t, "func Flag", moved)
}

func TestMoveBuildTagMismatchBlocked(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-build-tag/src"),
	)
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/tagged",
		[]string{"Flag"},
		"",
		"example/pkg/home",
		"existing.go",
		false,
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, "build constraint mismatch")
}

func TestMoveBuildTagMixedSourcesBlocked(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-build-tag/src"),
	)
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/tagged",
		[]string{"Flag", "Plain"},
		"",
		"example/pkg/tagged/hold",
		"hold.go",
		true,
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, "different build constraints")
}

func TestMoveBuildTagSameTagAppend(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-build-tag/src"),
	)
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/tagged",
		[]string{"Flag"},
		"",
		"example/pkg/home",
		"tagged.go",
		false,
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := service_tester.ReadFixtureFile(t, d, "pkg/home/tagged.go")
	assertFormatted(t, moved)
	assert.True(t, strings.HasPrefix(moved, "//go:build local"))
	assert.Integer(t, 1, strings.Count(moved, "//go:build"))
	assert.StringContains(t, "func Flag", moved)
	assert.StringContains(t, "func Local", moved)
}

func TestMoveBuildTagAbsentUnchanged(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-build-tag/src"),
	)
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/tagged",
		[]string{"Plain"},
		"",
		"example/pkg/tagged/hold",
		"hold.go",
		true,
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := service_tester.ReadFixtureFile(t, d, "pkg/tagged/hold/hold.go")
	assertFormatted(t, moved)
	assert.False(t, strings.Contains(moved, "//go:build"))
	assert.StringContains(t, "func Plain", moved)
}
