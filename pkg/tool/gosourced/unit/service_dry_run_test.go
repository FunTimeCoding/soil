package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"strings"
	"testing"
)

func TestDryRunLeavesTreeUntouched(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("batch-file/src"),
	)
	s := testService()
	source := service_tester.ReadFixtureFile(t, d, "pkg/target/constant.go")
	run := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	target := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/constant/constant.go",
	)
	r, e := s.MoveSymbols(
		d,
		"example/pkg/target",
		nil,
		"pkg/target/constant.go",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		source,
		service_tester.ReadFixtureFile(t, d, "pkg/target/constant.go"),
	)
	assert.String(
		t,
		run,
		service_tester.ReadFixtureFile(t, d, "pkg/target/run.go"),
	)
	assert.String(
		t,
		target,
		service_tester.ReadFixtureFile(t, d, "pkg/target/constant/constant.go"),
	)
}

func TestDryRunReportsWhatTheRealRunReports(t *testing.T) {
	planned := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("batch-file/src"),
	)
	s := testService()
	dry, e := s.MoveSymbols(
		planned,
		"example/pkg/target",
		nil,
		"pkg/target/constant.go",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
		true,
	)
	assert.FatalOnError(t, e)
	applied := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("batch-file/src"),
	)
	real, f := s.MoveSymbols(
		applied,
		"example/pkg/target",
		nil,
		"pkg/target/constant.go",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
		false,
	)
	assert.FatalOnError(t, f)
	assert.Strings(
		t,
		service_tester.ConcernText(real),
		service_tester.ConcernText(dry),
	)
}

func TestDryRunMarksConcernsPlanned(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("batch-file/src"),
	)
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/target",
		nil,
		"pkg/target/constant.go",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
		true,
	)
	assert.FatalOnError(t, e)
	assert.Greater(t, 0, len(r.Entries))

	for _, c := range r.Entries {
		assert.True(t, c.Planned)
		assert.False(t, c.Fixed)
	}
}

func TestDryRunRenameLeavesFilesUntouched(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("batch-enum/src"),
	)
	s := testService()
	before := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	r, e := s.Rename(d, "example/pkg/target", "Run", "Execute", "", true)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		before,
		service_tester.ReadFixtureFile(t, d, "pkg/target/run.go"),
	)
	assert.True(t, strings.Contains(before, "func Run()"))
}
