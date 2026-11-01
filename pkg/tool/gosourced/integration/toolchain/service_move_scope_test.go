package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestMoveBesideAPackageLevelNameAliases(t *testing.T) {
	d := moveScopeFixture(t)
	r, e := service_tester.NewService().MoveSymbol(
		d,
		"example/pkg/alfa",
		"Moving",
		"example/pkg/delta",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assertCompiles(t, d)
	assert.String(
		t,
		"package charlie\n\nimport example \"example/pkg/delta\"\n\nfunc Use() int {\n\treturn example.Moving()\n}\n",
		service_tester.ReadFixtureFile(t, d, "pkg/charlie/use.go"),
	)
}

func TestMoveUnderAShadowedImportRefuses(t *testing.T) {
	d := moveScopeFixture(t)
	r, e := service_tester.NewService().MoveSymbol(
		d,
		"example/pkg/alfa",
		"Shifting",
		"example/pkg/delta",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "delta is taken at")
	testutil.AssertBlockedContains(t, r, "pkg/echo/use.go:11")
	assertCompiles(t, d)
}
