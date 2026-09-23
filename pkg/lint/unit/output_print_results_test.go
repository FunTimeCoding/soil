package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"testing"
)

func TestPrintResultsBlockedReturnsTrue(t *testing.T) {
	entries := []*concern.Concern{
		concern.NewFile("test", "finding", "pkg/foo.go", false),
	}
	testutil.CaptureStdout(
		func() {
			assert.True(t, output.PrintResults(entries, false))
		},
	)
}

func TestPrintResultsAppliedReturnsFalse(t *testing.T) {
	entries := []*concern.Concern{
		concern.NewFile("test", "fixed", "pkg/foo.go", true),
	}
	testutil.CaptureStdout(
		func() {
			assert.False(t, output.PrintResults(entries, false))
		},
	)
}

func TestPrintResultsEmptyReturnsFalse(t *testing.T) {
	assert.False(t, output.PrintResults(nil, false))
}

func TestPrintResultsVerboseShowsAll(t *testing.T) {
	entries := []*concern.Concern{
		concern.NewFile("a", "applied fix", "pkg/foo.go", true),
		concern.NewFile("b", "blocked finding", "pkg/bar.go", false),
	}
	result := testutil.CaptureStdout(
		func() {
			output.PrintResults(entries, false)
		},
	)
	assert.String(
		t,
		"pkg/foo.go: applied fix (auto-fixed)\npkg/bar.go: blocked finding\n",
		result,
	)
}

func TestPrintResultsSummaryDeduplicatesApplied(t *testing.T) {
	entries := []*concern.Concern{
		concern.NewFile("a", "first", "pkg/foo.go", true),
		concern.NewFile("b", "second", "pkg/foo.go", true),
	}
	result := testutil.CaptureStdout(
		func() {
			output.PrintResults(entries, true)
		},
	)
	assert.String(t, "pkg/foo.go\n", result)
}

func TestPrintResultsSummaryShowsBlockedDetailed(t *testing.T) {
	entries := []*concern.Concern{
		concern.NewFile("a", "applied", "pkg/foo.go", true),
		concern.NewFile("b", "blocked", "pkg/bar.go", false),
	}
	result := testutil.CaptureStdout(
		func() {
			output.PrintResults(entries, true)
		},
	)
	assert.String(t, "pkg/foo.go\npkg/bar.go: blocked\n", result)
}

func TestPrintResultsLineLevel(t *testing.T) {
	entries := []*concern.Concern{
		concern.NewLine("test", "finding", "pkg/foo.go", 42, "", false),
	}
	result := testutil.CaptureStdout(
		func() {
			output.PrintResults(entries, false)
		},
	)
	assert.String(t, "pkg/foo.go:42: finding\n", result)
}
