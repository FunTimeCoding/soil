package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gobreak"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/integration/toolchain/break_tester"
	"path/filepath"
	"testing"
)

func TestRunReadsEachStatus(t *testing.T) {
	results, passed, e := gobreak.Run(
		break_tester.Fixture(t, constant.FixtureBreaks),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 4, passed)
	assert.Count(t, 10, results)
	assert.String(t, "killed", results[0].Status)
	assert.Strings(t, []string{"TestGradeHigh"}, results[0].Tests)
	assert.String(t, "panic", results[1].Status)
	assert.Strings(t, []string{"TestRatio"}, results[1].Tests)
	assert.String(t, "no coverage", results[2].Status)
	assert.String(t, "shape/shape.go:7 never ran", results[2].Detail)
	assert.String(t, "survived", results[3].Status)
	assert.String(t, "unmatched", results[4].Status)
	assert.StringContains(t, "search not found", results[4].Detail)
	assert.String(t, "compile error", results[5].Status)
	assert.String(
		t,
		"shape/shape.go:6:5: declared and not used: unused",
		results[5].Detail,
	)
	assert.String(t, "runtime error", results[6].Status)
	assert.String(t, "fixture/shape", results[6].Detail)
	assert.String(t, "survived", results[7].Status)
	assert.String(t, "survived", results[8].Status)
	assert.String(t, "unmatched", results[9].Status)
	assert.StringContains(t, "empty search", results[9].Detail)
}

func TestRunMarksMissedExpectations(t *testing.T) {
	results, _, e := gobreak.Run(
		break_tester.Fixture(t, constant.FixtureBreaks),
	)
	assert.FatalOnError(t, e)
	var missed []string

	for _, r := range results {
		if r.Missed() {
			missed = append(missed, r.Name)
		}
	}

	assert.Strings(
		t,
		[]string{"absent_text", "unused_local", "init_panic", "empty_find"},
		missed,
	)
}

func TestRunFailsOnTestsPatternMatchingNothing(t *testing.T) {
	_, _, e := gobreak.Run(
		break_tester.Fixture(t, constant.FixtureNoTestsBreaks),
	)
	assert.FatalNotNil(t, e)
	assert.String(t, "baseline: the tests pattern matches no test", e.Error())
}

func TestRunFailsOnUnknownExpect(t *testing.T) {
	_, _, e := gobreak.Run(
		break_tester.Fixture(t, constant.FixtureUnknownExpectBreaks),
	)
	assert.FatalNotNil(t, e)
	assert.String(
		t,
		`break high_bound: expect "red" is neither killed nor survived`,
		e.Error(),
	)
}

func TestRunFailsOnFailingBaseline(t *testing.T) {
	path := break_tester.Fixture(t, constant.FixtureBreaks)
	system_tester.New(t).Write(
		filepath.Dir(filepath.Dir(path)),
		"shape/failing_test.go",
		constant.FixtureFailingTest,
	)
	_, _, e := gobreak.Run(path)
	assert.FatalNotNil(t, e)
	assert.String(
		t,
		"baseline: tests already fail: TestAlreadyFails",
		e.Error(),
	)
}
