package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocheck"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/integration/check_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"path/filepath"
	"testing"
)

func TestAuditReadsConfigurationFromRootNotWorkingDirectory(t *testing.T) {
	root, e := filepath.EvalSymlinks(t.TempDir())
	assert.FatalOnError(t, e)
	testutil.WriteModFile(t, root, "example")
	system_tester.New(t).Write(
		root,
		".claude/settings.local.json",
		"{\"permissions\": {\"allow\": [\"mcp__alfa__bravo\"]}}\n",
	)
	system_tester.New(t).Write(
		root,
		"strata/tool/goaudit.yaml",
		"model_context:\n  alfa: pkg/tool/goalfad\n",
	)
	o := option.New()
	o.Root = root
	o.Index = t.TempDir()
	printed := testutil.CaptureStdout(
		func() { gocheck.Print(check_tester.New(t).Run(o)) },
	)
	assert.StringContains(t, "mapped service has no model_context/", printed)
}

func TestLaterStagesReportAfterEarlierStageFinds(t *testing.T) {
	root, e := filepath.EvalSymlinks(t.TempDir())
	assert.FatalOnError(t, e)
	testutil.WriteModFile(t, root, "example")
	system_tester.New(t).Write(
		root,
		"pkg/tool/charlie/charlie.go",
		"package charlie\n",
	)
	system_tester.New(t).Write(
		root,
		"cmd/charlie/main.go",
		"package main\n\nfunc main() {}\n",
	)
	system_tester.New(t).Write(
		root,
		"pkg/spacing/spacing.go",
		"package spacing\n\nfunc Alfa() int {\n\tbravo := 1\n\treturn bravo\n}\n",
	)
	system_tester.New(t).Write(
		root,
		"pkg/response/types.go",
		"package response\n\ntype LinkResponse struct {\n\tName string\n}\n\ntype Note struct {\n\tBody string\n}\n",
	)
	o := option.New()
	o.Root = root
	o.Index = t.TempDir()
	blocked := false
	printed := testutil.CaptureStdout(
		func() { blocked = gocheck.Print(check_tester.New(t).Run(o)) },
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
