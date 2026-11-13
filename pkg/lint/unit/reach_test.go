package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/reach"
	"github.com/funtimecoding/soil/pkg/lint/unit/transcript_tester"
	"github.com/funtimecoding/soil/pkg/source/warm_forest"
	"path/filepath"
	"testing"
)

func TestReachReadsEditToolPaths(t *testing.T) {
	edits, e := reach.ReadTranscript(
		transcript_tester.Write(
			t,
			transcript_tester.ToolUse(
				"alfa",
				"Edit",
				map[string]any{"file_path": "/repository/alfa/alfa.go"},
			),
			transcript_tester.ToolUse(
				"bravo",
				"Write",
				map[string]any{"file_path": "/repository/doc/bravo.md"},
			),
			transcript_tester.ToolUse(
				"charlie",
				"Read",
				map[string]any{"file_path": "/repository/charlie.go"},
			),
		),
		nil,
	)
	assert.FatalOnError(t, e)
	assert.Strings(
		t,
		[]string{"/repository/alfa/alfa.go", "/repository/doc/bravo.md"},
		edits.Files,
	)
}

func TestReachResolvesSourceResultsAgainstChosenModule(t *testing.T) {
	edits, e := reach.ReadTranscript(
		transcript_tester.Write(
			t,
			transcript_tester.ToolUse(
				"alfa",
				constant.SourceUseModule,
				map[string]any{"module": "delta"},
			),
			transcript_tester.ToolUse(
				"bravo",
				"mcp__source__rename_symbol",
				map[string]any{"symbol": "Echo"},
			),
			transcript_tester.ToolResult(
				"bravo",
				"pkg/echo/echo.go:3: Echo → Foxtrot\npkg/golf/golf.go:9: Echo → Foxtrot",
			),
		),
		map[string]string{"delta": "/delta"},
	)
	assert.FatalOnError(t, e)
	assert.Strings(
		t,
		[]string{"/delta/pkg/echo/echo.go", "/delta/pkg/golf/golf.go"},
		edits.Files,
	)
}

func TestReachSkipsDryRunSourceResults(t *testing.T) {
	edits, e := reach.ReadTranscript(
		transcript_tester.Write(
			t,
			transcript_tester.ToolUse(
				"alfa",
				"mcp__source__rename_symbol",
				map[string]any{"dry_run": true},
			),
			transcript_tester.ToolResult(
				"alfa",
				"/delta/echo.go:3: Echo → Golf",
			),
		),
		nil,
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, edits.Files)
}

func TestReachResolvesAndCountsBashWritesFromLineDirectory(t *testing.T) {
	edits, e := reach.ReadTranscript(
		transcript_tester.Write(
			t,
			transcript_tester.BashUse(
				"alfa",
				"/repository",
				"git mv alfa.go bravo/alfa.go",
			),
			transcript_tester.BashUse("bravo", "/repository", "ls -la"),
			transcript_tester.BashUse("charlie", "/repository", "rm $f"),
			transcript_tester.BashUse("delta", "/repository", "golint --fix"),
			transcript_tester.BashUse("echo", "/repository", "python3 run.py"),
			transcript_tester.BashUse("foxtrot", "/repository", "rm *.md"),
		),
		nil,
	)
	assert.FatalOnError(t, e)
	assert.Strings(
		t,
		[]string{"/repository/alfa.go", "/repository/bravo/alfa.go"},
		edits.Files,
	)
	assert.Strings(t, []string{"/repository/*.md"}, edits.Patterns)
	assert.Integer(t, 1, edits.Resolved)
	assert.Integer(t, 1, edits.Globbed)
	assert.Integer(t, 1, edits.Unresolved)
	assert.Integer(t, 1, edits.Fixers)
	assert.Integer(t, 1, edits.Scripts)
}

func TestReachTakesImportersAndKeepsDocumentsApart(t *testing.T) {
	root := t.TempDir()
	testutil.WriteModFile(t, root, "example.org/reach")
	testutil.WriteFile(
		t,
		root,
		"alfa/alfa.go",
		"package alfa\n\nfunc Alfa() {}\n",
	)
	testutil.WriteFile(
		t,
		root,
		"bravo/bravo.go",
		"package bravo\n\nimport \"example.org/reach/alfa\"\n\nfunc Bravo() { alfa.Alfa() }\n",
	)
	testutil.WriteFile(t, root, "charlie/charlie.go", "package charlie\n")
	testutil.WriteFile(t, root, "alfa/testdata/delta.go", "package delta\n")
	r := reach.Directories(
		warm_forest.New(),
		root,
		[]string{
			filepath.Join(root, "alfa/alfa.go"),
			filepath.Join(root, "alfa/testdata/delta.go"),
			filepath.Join(root, "doc/echo.md"),
		},
	)
	assert.Strings(t, []string{"./alfa", "./bravo"}, r.Packages)
	assert.Strings(t, []string{"./doc"}, r.Documents)
}

func TestReachKeepsImportersOfDeletedPackageAndDropsIt(t *testing.T) {
	root := t.TempDir()
	testutil.WriteModFile(t, root, "example.org/reach")
	testutil.WriteFile(
		t,
		root,
		"bravo/bravo.go",
		"package bravo\n\nimport \"example.org/reach/alfa\"\n\nfunc Bravo() { alfa.Alfa() }\n",
	)
	r := reach.Directories(
		warm_forest.New(),
		root,
		[]string{filepath.Join(root, "alfa/alfa.go")},
	)
	assert.Strings(t, []string{"./bravo"}, r.Packages)
}

func TestReachCarriesGraphItWidenedThrough(t *testing.T) {
	root := t.TempDir()
	testutil.WriteModFile(t, root, "example.org/reach")
	testutil.WriteFile(
		t,
		root,
		"alfa/alfa.go",
		"package alfa\n\nfunc Alfa() {}\n",
	)
	r := reach.Directories(
		warm_forest.New(),
		root,
		[]string{filepath.Join(root, "alfa/alfa.go")},
	)
	assert.FatalNotNil(t, r.Graph)
	_, okay := r.Graph.Nodes["example.org/reach/alfa"]
	assert.True(t, okay)
	documents := reach.Directories(
		warm_forest.New(),
		root,
		[]string{filepath.Join(root, "doc/bravo.md")},
	)
	assert.Nil(t, documents.Graph)
}
