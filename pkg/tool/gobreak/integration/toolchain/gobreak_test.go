package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gobreak"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
	"path/filepath"
	"testing"
)

func TestResultsReadEachKind(t *testing.T) {
	results, passed, e := gobreak.Results(fixture(t, constant.FixtureBreaks))
	assert.FatalOnError(t, e)
	assert.Integer(t, 4, passed)
	assert.Count(t, 10, results)
	assert.String(t, "red", results[0].Kind)
	assert.Strings(t, []string{"TestGradeHigh"}, results[0].Tests)
	assert.String(t, "panic", results[1].Kind)
	assert.Strings(t, []string{"TestRatio"}, results[1].Tests)
	assert.String(t, "unreached", results[2].Kind)
	assert.String(t, "shape/shape.go:7 never ran", results[2].Detail)
	assert.String(t, "survived", results[3].Kind)
	assert.String(t, "unmatched", results[4].Kind)
	assert.StringContains(t, "search not found", results[4].Detail)
	assert.String(t, "build failed", results[5].Kind)
	assert.String(
		t,
		"shape/shape.go:6:5: declared and not used: unused",
		results[5].Detail,
	)
	assert.String(t, "setup failed", results[6].Kind)
	assert.String(t, "fixture/shape", results[6].Detail)
	assert.String(t, "survived", results[7].Kind)
	assert.String(t, "survived", results[8].Kind)
	assert.String(t, "unmatched", results[9].Kind)
	assert.StringContains(t, "empty search", results[9].Detail)
}

func TestResultsMarkMissedExpectations(t *testing.T) {
	results, _, e := gobreak.Results(fixture(t, constant.FixtureBreaks))
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

func TestResultsRefuseTestsPatternMatchingNothing(t *testing.T) {
	_, _, e := gobreak.Results(fixture(t, constant.FixtureNoTestsBreaks))
	assert.FatalNotNil(t, e)
	assert.String(t, "baseline: the tests pattern matches no test", e.Error())
}

func TestResultsRefuseFailingBaseline(t *testing.T) {
	path := fixture(t, constant.FixtureBreaks)
	write(
		filepath.Dir(filepath.Dir(path)),
		"shape/failing_test.go",
		constant.FixtureFailingTest,
	)
	_, _, e := gobreak.Results(path)
	assert.FatalNotNil(t, e)
	assert.String(
		t,
		"baseline: tests already fail: TestAlreadyFails",
		e.Error(),
	)
}
