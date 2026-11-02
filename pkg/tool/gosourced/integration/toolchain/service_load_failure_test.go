package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/load_failure"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestABrokenCallerNamesItself(t *testing.T) {
	d := renameFunctionFixture(t)
	testutil.WriteFile(
		t,
		d,
		"pkg/broken/broken.go",
		"package broken\n\nimport \"example/pkg/target\"\n\nfunc Broken() bool {\n\treturn target.IsGeneratedHeader(missing)\n}\n",
	)
	_, e := service_tester.NewService(t).Rename(
		t.Context(),
		d,
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

func TestABrokenCallerNobodyTouchesIsSettled(t *testing.T) {
	d := renameFunctionFixture(t)
	testutil.WriteFile(
		t,
		d,
		"pkg/broken/broken.go",
		"package broken\n\nimport \"example/pkg/target\"\n\nfunc Broken() bool {\n\treturn target.IsGeneratedHeader(missing)\n}\n",
	)
	hourAgo := time.Now().Add(-time.Hour)
	assert.FatalOnError(
		t,
		os.Chtimes(filepath.Join(d, "pkg/broken/broken.go"), hourAgo, hourAgo),
	)
	_, e := service_tester.NewService(t).Rename(
		t.Context(),
		d,
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

func TestABrokenPackageOutsideTheReachDoesNotBlock(t *testing.T) {
	d := renameFunctionFixture(t)
	testutil.WriteFile(
		t,
		d,
		"pkg/broken/broken.go",
		"package broken\n\nfunc Broken() int {\n\treturn missing\n}\n",
	)
	r, e := service_tester.NewService(t).Rename(
		t.Context(),
		d,
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
		service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go"),
	)
}

func TestABrokenPackageOutsideTheReachDoesNotBlockAWholeLoad(t *testing.T) {
	d := renameFunctionFixture(t)
	testutil.WriteFile(
		t,
		d,
		"pkg/broken/broken.go",
		"package broken\n\nfunc Broken() int {\n\treturn missing\n}\n",
	)
	s := service_tester.NewService(t)
	s.UseFullLoad()
	r, e := s.Rename(
		t.Context(),
		d,
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
		service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go"),
	)
}

func TestABrokenPackageOutsideTheReachInAReplacingModuleDoesNotBlock(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "rename-replacer")
	testutil.WriteFile(
		t,
		user,
		"pkg/broken/broken.go",
		"package broken\n\nfunc Broken() int {\n\treturn missing\n}\n",
	)
	r, e := replacingService(t, library, user).Rename(
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
	assert.StringContains(
		t,
		"lib.Write(\"a\")",
		service_tester.ReadFixtureFile(t, user, "pkg/user/user.go"),
	)
}
