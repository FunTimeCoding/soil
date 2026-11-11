package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestMoveRenamesShadowingLocal(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-shadowing-local/src"),
	)
	r, e := service_tester.NewService(t).Symbol.MoveSymbol(
		t.Context(),
		d,
		"example/pkg/alfa/bravo",
		"Limit",
		"example/pkg/alfa/echo",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"package delta\n\nimport \"example/pkg/alfa/echo\"\n\nfunc Run() int {\n\te := echo.Limit()\n\n\treturn e * 2\n}\n",
		service_tester.ReadFixtureFile(t, d, "pkg/alfa/delta/run.go"),
	)
	assert.String(
		t,
		"package golf\n\nimport \"example/pkg/alfa/echo\"\n\nfunc Run(e int) int {\n\tc := echo.Limit()\n\n\treturn c * e\n}\n",
		service_tester.ReadFixtureFile(t, d, "pkg/alfa/golf/run.go"),
	)
	assert.String(
		t,
		"package india\n\nimport \"example/pkg/alfa/echo\"\n\nfunc Run() int {\n\tc := echo.Limit()\n\n\tfor e := range 2 {\n\t\tc += e\n\t}\n\n\treturn c\n}\n",
		service_tester.ReadFixtureFile(t, d, "pkg/alfa/india/run.go"),
	)
}

func TestMoveRefusesShadowingLocalWithoutFreeLetter(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-shadowing-local-refused/src"),
	)
	r, e := service_tester.NewService(t).Symbol.MoveSymbol(
		t.Context(),
		d,
		"example/pkg/alfa/bravo",
		"Limit",
		"example/pkg/alfa/echo",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, "local echo")
	assert.String(
		t,
		"package hotel\n\nimport \"example/pkg/alfa/bravo\"\n\nfunc Run(\n\te int,\n\tc int,\n\th int,\n\to int,\n) int {\n\techo := bravo.Limit()\n\n\treturn echo * e * c * h * o\n}\n",
		service_tester.ReadFixtureFile(t, d, "pkg/alfa/hotel/run.go"),
	)
}
