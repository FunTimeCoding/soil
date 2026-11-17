package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestMoveBesidePackageLevelNameAliases(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-scope/src"))
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"Moving",
		"example/pkg/delta",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	rewrite_tester.Compiles(t, o.Directory)
	assert.String(
		t,
		"package charlie\n\nimport exampleDelta \"example/pkg/delta\"\n\nfunc Use() int {\n\treturn exampleDelta.Moving()\n}\n",
		o.Read("pkg/charlie/use.go"),
	)
}

func TestMoveUnderShadowedImportRefuses(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-scope/src"))
	r, e := o.Symbol.MoveSymbol(
		t.Context(),
		o.Directory,
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
	rewrite_tester.Compiles(t, o.Directory)
}
