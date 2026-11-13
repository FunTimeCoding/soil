package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze"
	goanalyzeOption "github.com/funtimecoding/soil/pkg/tool/goanalyze/option"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	gofixOption "github.com/funtimecoding/soil/pkg/tool/gofix/option"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
	"path/filepath"
	"testing"
)

func TestAnalysisOverCleanFixLoadMatchesFreshAnalysis(t *testing.T) {
	root, e := filepath.EvalSymlinks(t.TempDir())
	assert.FatalOnError(t, e)
	testutil.WriteModFile(t, root, "example")
	testutil.WriteFile(
		t,
		root,
		"pkg/response/types.go",
		"package response\n\ntype LinkResponse struct {\n\tName string\n}\n\ntype Note struct {\n\tBody string\n}\n",
	)
	patterns := []string{"./..."}
	directory := t.TempDir()
	reported, f := resolve.ListPackages(root, patterns...)
	assert.FatalOnError(t, f)
	handle := index.New(directory, root, goanalyze.Kinds()...)
	fix := gofixOption.New()
	fix.Root = root
	fix.Patterns = patterns
	fix.Diff = true
	fix.Index = directory
	fix.Reported = reported
	fix.Handle = handle
	w, failed := gofix.RunDefault(fix, output.NewResultsWithDirectory(root))
	assert.FatalOnError(t, failed)
	loaded, okay := gofix.RememberedLoad(w, root, patterns)
	assert.True(t, okay)
	analyze := goanalyzeOption.New()
	analyze.Root = root
	analyze.Patterns = patterns
	analyze.Index = directory
	shared := goanalyze.AnalyzeLoaded(analyze, loaded, reported, handle)
	fresh := goanalyze.Analyze(analyze)
	testutil.AssertBlockedContains(
		t,
		shared,
		"struct LinkResponse repeats the bag's name response - rename it Link",
	)
	assert.Integer(t, len(fresh.Entries), len(shared.Entries))
}

func TestFixThatRewritesLeavesNoLoadToReuse(t *testing.T) {
	root, e := filepath.EvalSymlinks(t.TempDir())
	assert.FatalOnError(t, e)
	testutil.WriteModFile(t, root, "example")
	testutil.WriteFile(
		t,
		root,
		"pkg/format/format.go",
		"package format\n\nimport \"fmt\"\n\nfunc Alfa() {\n\tfmt.Println(\"bravo charlie delta echo foxtrot\", \"golf hotel india juliett\", \"kilo\")\n}\n",
	)
	fix := gofixOption.New()
	fix.Root = root
	fix.Patterns = []string{"./..."}
	fix.Diff = true
	fix.Index = t.TempDir()
	var w *workspace.Workspace
	var f error
	testutil.CaptureStdout(
		func() {
			w, f = gofix.RunDefault(fix, output.NewResultsWithDirectory(root))
		},
	)
	assert.FatalOnError(t, f)
	_, okay := gofix.RememberedLoad(w, root, fix.Patterns)
	assert.False(t, okay)
}
