package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"io"
	"os"
	"testing"
)

func TestPrintResultsBlockedReturnsTrue(t *testing.T) {
	entries := []*concern.Concern{
		concern.NewFile("test", "finding", "pkg/foo.go", false),
	}
	captureStdout(
		func() {
			assert.True(t, output.PrintResults(entries, false))
		},
	)
}

func TestPrintResultsAppliedReturnsFalse(t *testing.T) {
	entries := []*concern.Concern{
		concern.NewFile("test", "fixed", "pkg/foo.go", true),
	}
	captureStdout(
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
	result := captureStdout(
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
	result := captureStdout(
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
	result := captureStdout(
		func() {
			output.PrintResults(entries, true)
		},
	)
	assert.String(
		t,
		"pkg/foo.go\npkg/bar.go: blocked\n",
		result,
	)
}

func TestPrintResultsLineLevel(t *testing.T) {
	entries := []*concern.Concern{
		concern.NewLine("test", "finding", "pkg/foo.go", 42, "", false),
	}
	result := captureStdout(
		func() {
			output.PrintResults(entries, false)
		},
	)
	assert.String(t, "pkg/foo.go:42: finding\n", result)
}

func captureStdout(f func()) string {
	original := os.Stdout
	reader, writer, e := os.Pipe()
	errors.PanicOnError(e)
	os.Stdout = writer
	f()
	errors.PanicClose(writer)
	os.Stdout = original
	captured, e := io.ReadAll(reader)
	errors.PanicOnError(e)

	return string(captured)
}
