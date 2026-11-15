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

func TestPackageThatDoesNotLoadIsFindingNotExit(t *testing.T) {
	root, e := filepath.EvalSymlinks(t.TempDir())
	assert.FatalOnError(t, e)
	testutil.WriteModFile(t, root, "example")
	system_tester.Write(
		t,
		root,
		"pkg/tool/charlie/charlie.go",
		"package charlie\n",
	)
	system_tester.Write(
		t,
		root,
		"cmd/charlie/main.go",
		"package main\n\nfunc main() {}\n",
	)
	system_tester.Write(
		t,
		root,
		"pkg/spacing/spacing.go",
		"package spacing\n\nfunc Alfa() int {\n\tbravo := 1\n\treturn bravo\n}\n",
	)
	system_tester.Write(
		t,
		root,
		"pkg/broken/broken.go",
		"package broken\n\nfunc Delta() int {\n\treturn echo\n}\n",
	)
	o := option.New()
	o.Root = root
	o.Index = t.TempDir()
	blocked := false
	printed := testutil.CaptureStdout(
		func() { blocked = gocheck.Print(check_tester.Run(t, o)) },
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
		"pkg/broken/broken.go:4:9: does not load: example/pkg/broken - undefined: echo",
		printed,
	)
}
