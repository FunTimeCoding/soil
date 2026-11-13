package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/load_failure"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gocheck"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	goclaudeConstant "github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	gofixOption "github.com/funtimecoding/soil/pkg/tool/gofix/option"
	"path/filepath"
	"testing"
)

func TestRunPrintsNothing(t *testing.T) {
	root := stackModule(t)
	o := option.New()
	o.Root = root
	o.Index = t.TempDir()
	printed := testutil.CaptureStdout(func() { checked(t, o) })
	assert.String(t, "", printed)
}

func TestReachWithoutSessionIsErrorNotExit(t *testing.T) {
	t.Setenv(goclaudeConstant.SessionIdentifierEnvironment, "")
	_, e := gocheck.Transcript(constant.OwnSession)
	assert.True(t, not_selected.Is(e))
}

func TestFixReturnsLoadFailureInsteadOfExiting(t *testing.T) {
	root, e := filepath.EvalSymlinks(t.TempDir())
	assert.FatalOnError(t, e)
	testutil.WriteModFile(t, root, "example")
	testutil.WriteFile(
		t,
		root,
		"pkg/broken/broken.go",
		"package broken\n\nfunc Delta() int {\n\treturn echo\n}\n",
	)
	fix := gofixOption.New()
	fix.Root = root
	fix.Patterns = []string{"./..."}
	fix.Diff = true
	fix.Index = t.TempDir()
	_, failed := gofix.RunDefault(fix, output.NewResultsWithDirectory(root))
	failure, okay := load_failure.As(failed)
	assert.True(t, okay)
	assert.String(t, "example/pkg/broken", failure.Package)
}
