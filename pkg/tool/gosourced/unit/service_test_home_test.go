package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"os"
	"path/filepath"
	"testing"
)

func TestRenamePackageTestHome(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("test-home/src"))
	s := testService()
	r, e := s.RenamePackage(d, "example/pkg/alfa/unit_test", "unit", false)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	_, e = os.Stat(filepath.Join(d, "pkg/alfa/unit_test"))
	assert.True(t, os.IsNotExist(e))
	moved := readFixtureFile(t, d, "pkg/alfa/unit/parse_test.go")
	assert.StringContains(t, "package unit\n", moved)
	assert.StringContains(t, "example/pkg/alfa/unit/alfa_tester", moved)
	tester := readFixtureFile(t, d, "pkg/alfa/unit/alfa_tester/sample.go")
	assert.StringContains(t, "package alfa_tester", tester)
}

func TestMovePackageTestFacet(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("test-home/src"))
	e := os.MkdirAll(filepath.Join(d, "pkg/alfa/integration"), 0755)
	assert.FatalOnError(t, e)
	s := testService()
	r, f := s.MovePackage(
		d,
		"example/pkg/alfa/integration_test/client",
		"example/pkg/alfa/integration/client",
		false,
	)
	assert.FatalOnError(t, f)
	testutil.AssertBlocked(t, r, 0)
	_, f = os.Stat(filepath.Join(d, "pkg/alfa/integration_test/client"))
	assert.True(t, os.IsNotExist(f))
	moved := readFixtureFile(t, d, "pkg/alfa/integration/client/client_test.go")
	assert.StringContains(t, "package client", moved)
}
