package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestRenamePackageTestHome(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("test-home/src"))
	r, e := o.Packages.RenamePackage(
		t.Context(),
		o.Directory,
		"example/pkg/alfa/unit_test",
		constant.UnitSegment,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	_, e = os.Stat(filepath.Join(o.Directory, "pkg/alfa/unit_test"))
	assert.True(t, os.IsNotExist(e))
	moved := o.Read("pkg/alfa/unit/parse_test.go")
	assert.StringContains(t, "package unit\n", moved)
	assert.StringContains(t, "example/pkg/alfa/unit/alfa_tester", moved)
	tester := o.Read("pkg/alfa/unit/alfa_tester/sample.go")
	assert.StringContains(t, "package alfa_tester", tester)
}

func TestMovePackageTestFacet(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("test-home/src"))
	e := os.MkdirAll(filepath.Join(o.Directory, "pkg/alfa/integration"), 0755)
	assert.FatalOnError(t, e)
	r, f := o.Packages.MovePackage(
		t.Context(),
		o.Directory,
		"example/pkg/alfa/integration_test/client",
		"example/pkg/alfa/integration/client",
		false,
	)
	assert.FatalOnError(t, f)
	testutil.AssertBlocked(t, r, 0)
	_, f = os.Stat(
		filepath.Join(o.Directory, "pkg/alfa/integration_test/client"),
	)
	assert.True(t, os.IsNotExist(f))
	moved := o.Read("pkg/alfa/integration/client/client_test.go")
	assert.StringContains(t, "package client", moved)
}
