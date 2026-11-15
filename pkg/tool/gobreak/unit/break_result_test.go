package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/break_result"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
	"testing"
)

func TestBreakResultLineNamesTests(t *testing.T) {
	r := break_result.New(
		"high_bound",
		constant.StatusKilled,
		constant.StatusKilled,
	)
	r.Tests = []string{"TestGradeHigh", "TestGradeLow"}
	assert.String(t, "high_bound  killed  TestGradeHigh TestGradeLow", r.Line())
}

func TestBreakResultLineMarksMissedExpectation(t *testing.T) {
	r := break_result.New(
		"absent_text",
		constant.StatusKilled,
		constant.StatusUnmatched,
	)
	r.Detail = "shape.go: block 1: search not found"
	assert.String(
		t,
		"absent_text  unmatched  shape.go: block 1: search not found  (expected killed)",
		r.Line(),
	)
}
