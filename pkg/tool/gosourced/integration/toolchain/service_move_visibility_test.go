package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestMoveKeepsHelperOnlyBatchCallsPrivate(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-visibility/src"),
	)
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		[]string{"Build", "piece"},
		"",
		"example/pkg/bravo",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	rewrite_tester.Compiles(t, o.Directory)
	assert.StringContains(t, "func piece() int", o.Read("pkg/bravo/piece.go"))
	assert.StringContains(t, "return piece() + 1", o.Read("pkg/bravo/build.go"))
}

func TestMoveExportsWhatStaysBehindCalls(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-visibility/src"),
	)
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		[]string{"lone"},
		"",
		"example/pkg/bravo",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	rewrite_tester.Compiles(t, o.Directory)
	assert.StringContains(t, "func Lone() int", o.Read("pkg/bravo/lone.go"))
	assert.StringContains(t, "bravo.Lone()", o.Read("pkg/alfa/keep.go"))
}

func TestMoveExportsWhatInPackageTestCalls(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-visibility/src"),
	)
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		[]string{"probe"},
		"",
		"example/pkg/bravo",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	rewrite_tester.Compiles(t, o.Directory)
	assert.StringContains(t, "func Probe() int", o.Read("pkg/bravo/probe.go"))
}
