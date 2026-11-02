package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestMoveKeepsAHelperOnlyTheBatchCallsPrivate(t *testing.T) {
	d := moveVisibilityFixture(t)
	r, e := service_tester.NewService(t).MoveSymbols(
		t.Context(),
		d,
		"example/pkg/alfa",
		[]string{"Build", "piece"},
		"",
		"example/pkg/bravo",
		"",
		false,
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assertCompiles(t, d)
	assert.StringContains(
		t,
		"func piece() int",
		service_tester.ReadFixtureFile(t, d, "pkg/bravo/piece.go"),
	)
	assert.StringContains(
		t,
		"return piece() + 1",
		service_tester.ReadFixtureFile(t, d, "pkg/bravo/build.go"),
	)
}

func TestMoveExportsWhatStaysBehindCalls(t *testing.T) {
	d := moveVisibilityFixture(t)
	r, e := service_tester.NewService(t).MoveSymbols(
		t.Context(),
		d,
		"example/pkg/alfa",
		[]string{"lone"},
		"",
		"example/pkg/bravo",
		"",
		false,
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assertCompiles(t, d)
	assert.StringContains(
		t,
		"func Lone() int",
		service_tester.ReadFixtureFile(t, d, "pkg/bravo/lone.go"),
	)
	assert.StringContains(
		t,
		"bravo.Lone()",
		service_tester.ReadFixtureFile(t, d, "pkg/alfa/keep.go"),
	)
}

func TestMoveExportsWhatAnInPackageTestCalls(t *testing.T) {
	d := moveVisibilityFixture(t)
	r, e := service_tester.NewService(t).MoveSymbols(
		t.Context(),
		d,
		"example/pkg/alfa",
		[]string{"probe"},
		"",
		"example/pkg/bravo",
		"",
		false,
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assertCompiles(t, d)
	assert.StringContains(
		t,
		"func Probe() int",
		service_tester.ReadFixtureFile(t, d, "pkg/bravo/probe.go"),
	)
}
