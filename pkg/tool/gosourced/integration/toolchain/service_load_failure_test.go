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

func TestABrokenPackageElsewhereNamesItself(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-function/src"),
	)
	testutil.WriteFile(
		t,
		d,
		"pkg/broken/broken.go",
		"package broken\n\nfunc Broken() int {\n\treturn missing\n}\n",
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

func TestABrokenPackageNobodyTouchesIsSettled(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-function/src"),
	)
	testutil.WriteFile(
		t,
		d,
		"pkg/broken/broken.go",
		"package broken\n\nfunc Broken() int {\n\treturn missing\n}\n",
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
