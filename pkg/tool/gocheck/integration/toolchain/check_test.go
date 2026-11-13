package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gocheck"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"path/filepath"
	"testing"
)

func TestAuditReadsConfigurationFromRootNotWorkingDirectory(t *testing.T) {
	root, e := filepath.EvalSymlinks(t.TempDir())
	assert.FatalOnError(t, e)
	testutil.WriteModFile(t, root, "example")
	testutil.WriteFile(
		t,
		root,
		".claude/settings.local.json",
		"{\"permissions\": {\"allow\": [\"mcp__alfa__bravo\"]}}\n",
	)
	testutil.WriteFile(
		t,
		root,
		"strata/tool/goaudit.yaml",
		"model_context:\n  alfa: pkg/tool/goalfad\n",
	)
	o := option.New()
	o.Root = root
	o.Index = t.TempDir()
	printed := testutil.CaptureStdout(func() { gocheck.Print(checked(t, o)) })
	assert.StringContains(t, "mapped service has no model_context/", printed)
}

func TestLaterStagesReportAfterEarlierStageFinds(t *testing.T) {
	root, e := filepath.EvalSymlinks(t.TempDir())
	assert.FatalOnError(t, e)
	testutil.WriteModFile(t, root, "example")
	testutil.WriteFile(
		t,
		root,
		"pkg/tool/charlie/charlie.go",
		"package charlie\n",
	)
	testutil.WriteFile(
		t,
		root,
		"cmd/charlie/main.go",
		"package main\n\nfunc main() {}\n",
	)
	testutil.WriteFile(
		t,
		root,
		"pkg/spacing/spacing.go",
		"package spacing\n\nfunc Alfa() int {\n\tbravo := 1\n\treturn bravo\n}\n",
	)
	testutil.WriteFile(
		t,
		root,
		"pkg/response/types.go",
		"package response\n\ntype LinkResponse struct {\n\tName string\n}\n\ntype Note struct {\n\tBody string\n}\n",
	)
	o := option.New()
	o.Root = root
	o.Index = t.TempDir()
	blocked := false
	printed := testutil.CaptureStdout(
		func() { blocked = gocheck.Print(checked(t, o)) },
	)
	assert.True(t, blocked)
	assert.StringContains(
		t,
		"pkg/spacing/spacing.go:5: Missing blank line before return/break/continue",
		printed,
	)
	assert.StringContains(t, "cmd/charlie: missing sentry reporter", printed)
	assert.StringContains(
		t,
		"pkg/response/types.go:3: struct LinkResponse repeats the bag's name response - rename it Link",
		printed,
	)
}
