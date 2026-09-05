package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"strings"
	"testing"
)

func concernText(r *output.Results) []string {
	var result []string

	for _, c := range r.Entries {
		result = append(result, fmt.Sprintf("%s: %s", c.Path, c.Text))
	}

	return result
}

func TestDryRunLeavesTreeUntouched(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("batch-file/src"))
	s := testService()
	source := readFixtureFile(t, d, "pkg/target/constant.go")
	run := readFixtureFile(t, d, "pkg/target/run.go")
	target := readFixtureFile(t, d, "pkg/target/constant/constant.go")
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
	assert.String(t, source, readFixtureFile(t, d, "pkg/target/constant.go"))
	assert.String(t, run, readFixtureFile(t, d, "pkg/target/run.go"))
	assert.String(
		t,
		target,
		readFixtureFile(t, d, "pkg/target/constant/constant.go"),
	)
}

func TestDryRunReportsWhatTheRealRunReports(t *testing.T) {
	planned := testutil.PrepareTestPackage(t, serviceTestdata("batch-file/src"))
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
	applied := testutil.PrepareTestPackage(t, serviceTestdata("batch-file/src"))
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
	assert.Strings(t, concernText(real), concernText(dry))
}

func TestDryRunMarksConcernsPlanned(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("batch-file/src"))
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
	d := testutil.PrepareTestPackage(t, serviceTestdata("batch-enum/src"))
	s := testService()
	before := readFixtureFile(t, d, "pkg/target/run.go")
	r, e := s.Rename(d, "example/pkg/target", "Run", "Execute", "", true)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(t, before, readFixtureFile(t, d, "pkg/target/run.go"))
	assert.True(t, strings.Contains(before, "func Run()"))
}
