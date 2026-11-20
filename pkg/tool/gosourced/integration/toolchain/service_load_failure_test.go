package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/load_failure"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/module_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestBrokenCallerNamesItself(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-function/src"),
	)
	system_tester.New(t).Write(
		o.Directory,
		"pkg/broken/broken.go",
		"package broken\n\nimport \"example/pkg/target\"\n\nfunc Broken() bool {\n\treturn target.IsGeneratedHeader(missing)\n}\n",
	)
	_, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"IsGeneratedHeader",
		"IsGenerated",
		"",
		false,
	)
	failure, okay := load_failure.As(e)
	assert.True(t, okay)
	assert.String(t, "example/pkg/broken", failure.Package)
	assert.False(t, failure.Settled)
}

func TestBrokenCallerNobodyTouchesIsSettled(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-function/src"),
	)
	system_tester.New(t).Write(
		o.Directory,
		"pkg/broken/broken.go",
		"package broken\n\nimport \"example/pkg/target\"\n\nfunc Broken() bool {\n\treturn target.IsGeneratedHeader(missing)\n}\n",
	)
	hourAgo := time.Now().Add(-time.Hour)
	assert.FatalOnError(
		t,
		os.Chtimes(
			filepath.Join(o.Directory, "pkg/broken/broken.go"),
			hourAgo,
			hourAgo,
		),
	)
	_, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"IsGeneratedHeader",
		"IsGenerated",
		"",
		false,
	)
	failure, okay := load_failure.As(e)
	assert.True(t, okay)
	assert.True(t, failure.Settled)
}

func TestBrokenPackageOutsideReachDoesNotBlock(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-function/src"),
	)
	system_tester.New(t).Write(
		o.Directory,
		"pkg/broken/broken.go",
		"package broken\n\nfunc Broken() int {\n\treturn missing\n}\n",
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"IsGeneratedHeader",
		"IsGenerated",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(
		t,
		"target.IsGenerated(content)",
		o.Read("pkg/caller/run.go"),
	)
}

func TestBrokenPackageOutsideReachDoesNotBlockWholeLoad(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-function/src"),
	)
	system_tester.New(t).Write(
		o.Directory,
		"pkg/broken/broken.go",
		"package broken\n\nfunc Broken() int {\n\treturn missing\n}\n",
	)
	o.Workspace.UseFullLoad()
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"IsGeneratedHeader",
		"IsGenerated",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(
		t,
		"target.IsGenerated(content)",
		o.Read("pkg/caller/run.go"),
	)
}

func TestBrokenPackageOutsideReachInReplacingModuleDoesNotBlock(t *testing.T) {
	library, o := module_tester.New(t).ReplacedLibrary("rename-replacer")
	system_tester.New(t).Write(
		o.Directory,
		"pkg/broken/broken.go",
		"package broken\n\nfunc Broken() int {\n\treturn missing\n}\n",
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		library,
		"other.test/lib",
		"Log",
		"Write",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(t, "lib.Write(\"a\")", o.Read("pkg/user/user.go"))
}
