package gofix

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"io"
	"os"
	"testing"
)

func TestPrintResults(t *testing.T) {
	t.Run(
		"VerboseApplied",
		func(t *testing.T) {
			entries := []*concern.Concern{
				concern.NewFile(
					"import_alias",
					"de-aliased sentry → constant",
					"pkg/foo/bar.go",
					true,
				),
				concern.NewFile(
					"renamed",
					"renamed dirName → directoryName (4 references)",
					"pkg/foo/baz.go",
					true,
				),
				concern.NewFile(
					"call_format",
					"formatted call (line 42)",
					"pkg/foo/baz.go",
					true,
				),
			}
			captured := captureStdout(
				func() {
					hasBlocked := output.PrintResults(entries, false)
					assert.False(t, hasBlocked)
				},
			)
			assert.String(
				t,
				"pkg/foo/bar.go: de-aliased sentry → constant (auto-fixed)\npkg/foo/baz.go: renamed dirName → directoryName (4 references) (auto-fixed)\npkg/foo/baz.go: formatted call (line 42) (auto-fixed)\n",
				captured,
			)
		},
	)
	t.Run(
		"SummaryDeduplicates",
		func(t *testing.T) {
			entries := []*concern.Concern{
				concern.NewFile(
					"import_alias",
					"de-aliased sentry → constant",
					"pkg/foo/bar.go",
					true,
				),
				concern.NewFile(
					"renamed",
					"renamed dirName → directoryName (4 references)",
					"pkg/foo/baz.go",
					true,
				),
				concern.NewFile(
					"call_format",
					"formatted call (line 42)",
					"pkg/foo/baz.go",
					true,
				),
			}
			captured := captureStdout(
				func() {
					hasBlocked := output.PrintResults(entries, true)
					assert.False(t, hasBlocked)
				},
			)
			assert.String(
				t,
				"pkg/foo/bar.go\npkg/foo/baz.go\n",
				captured,
			)
		},
	)
	t.Run(
		"BlockedAlwaysDetailed",
		func(t *testing.T) {
			entries := []*concern.Concern{
				concern.NewFile(
					"import_alias",
					"de-aliased h → helper",
					"pkg/foo/bar.go",
					true,
				),
				concern.NewFile(
					"import_alias",
					"cannot de-alias mark → server (local collision with \"server\")",
					"pkg/foo/baz.go",
					false,
				),
			}
			captured := captureStdout(
				func() {
					hasBlocked := output.PrintResults(entries, true)
					assert.True(t, hasBlocked)
				},
			)
			assert.String(
				t,
				"pkg/foo/bar.go\npkg/foo/baz.go: cannot de-alias mark → server (local collision with \"server\")\n",
				captured,
			)
		},
	)
	t.Run(
		"BlockedVerbose",
		func(t *testing.T) {
			entries := []*concern.Concern{
				concern.NewFile(
					"import_alias",
					"de-aliased h → helper",
					"pkg/foo/bar.go",
					true,
				),
				concern.NewFile(
					"collision",
					"cannot rename dirName (collision)",
					"pkg/foo/baz.go",
					false,
				),
			}
			captured := captureStdout(
				func() {
					hasBlocked := output.PrintResults(entries, false)
					assert.True(t, hasBlocked)
				},
			)
			assert.String(
				t,
				"pkg/foo/bar.go: de-aliased h → helper (auto-fixed)\npkg/foo/baz.go: cannot rename dirName (collision)\n",
				captured,
			)
		},
	)
	t.Run(
		"EmptyNoOutput",
		func(t *testing.T) {
			captured := captureStdout(
				func() {
					hasBlocked := output.PrintResults(nil, false)
					assert.False(t, hasBlocked)
				},
			)
			assert.String(t, "", captured)
		},
	)
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
