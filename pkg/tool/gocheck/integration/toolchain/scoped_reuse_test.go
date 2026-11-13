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
	"path/filepath"
	"testing"
)

func TestScopedLoadHandedToFixServesScopedAnalysis(t *testing.T) {
	root, e := filepath.EvalSymlinks(t.TempDir())
	assert.FatalOnError(t, e)
	testutil.WriteModFile(t, root, "example")
	testutil.WriteFile(
		t,
		root,
		"pkg/response/types.go",
		"package response\n\ntype LinkResponse struct {\n\tName string\n}\n\ntype Note struct {\n\tBody string\n}\n",
	)
	testutil.WriteFile(
		t,
		root,
		"pkg/other/other.go",
		"package other\n\ntype FoxtrotOther struct {\n\tName string\n}\n\ntype Golf struct {\n\tBody string\n}\n",
	)
	patterns := []string{"./pkg/response"}
	directory := t.TempDir()
	reported, f := resolve.ListPackages(root, patterns...)
	assert.FatalOnError(t, f)
	loaded, set, g := resolve.LoadPackagesUnchecked(root, nil, patterns...)
	assert.FatalOnError(t, g)
	handle := index.New(directory, root, goanalyze.Kinds()...)
	fix := gofixOption.New()
	fix.Root = root
	fix.Patterns = patterns
	fix.Diff = true
	fix.Index = directory
	fix.Reported = reported
	fix.Handle = handle
	fix.Loaded = loaded
	fix.FileSet = set
	w, failed := gofix.RunDefault(fix, output.NewResultsWithDirectory(root))
	assert.FatalOnError(t, failed)
	remembered, okay := gofix.RememberedLoad(w, root, patterns)
	assert.True(t, okay)
	assert.True(t, remembered[0] == loaded[0])
	analyze := goanalyzeOption.New()
	analyze.Root = root
	analyze.Patterns = patterns
	analyze.Index = directory
	shared := goanalyze.AnalyzeLoaded(analyze, remembered, reported, handle)
	fresh := goanalyze.Analyze(analyze)
	testutil.AssertBlockedContains(
		t,
		shared,
		"struct LinkResponse repeats the bag's name response - rename it Link",
	)
	testutil.AssertNotBlockedContains(t, shared, "FoxtrotOther")
	assert.Integer(t, len(fresh.Entries), len(shared.Entries))
}
