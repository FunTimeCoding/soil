package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestApplyPatternRewritesCallWhereverItSits(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("pattern-expression/src"),
	)
	r, apply, e := service_tester.NewService(t).Pattern.ApplyPattern(
		t.Context(),
		d,
		"example/pkg/alfa",
		"New",
		"",
		"func pattern(n string) {\n\talfa.New(n)\n}",
		"func replacement(n string) {\n\talfa.Named(n)\n}",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.Integer(t, 4, apply.Total)
	assert.Integer(t, 4, apply.Matched)
	assert.Integer(t, 4, apply.Rewritten)
	holder := service_tester.ReadFixtureFile(t, d, "pkg/bravo/holder.go")
	assert.StringContains(t, "&Holder{Server: alfa.Named(\"bravo\")}", holder)
	assert.StringContains(t, "describe(alfa.Named(\"charlie\"))", holder)
	assert.StringContains(t, "\talfa.Named(\"delta\")\n", holder)
	assert.StringContains(t, "return alfa.Named(\"echo\")", holder)
	assertCompiles(t, d)
}

func TestApplyPatternRefusesNestedCalls(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("pattern-expression/src"),
	)
	r, apply, e := service_tester.NewService(t).Pattern.ApplyPattern(
		t.Context(),
		d,
		"example/pkg/alfa",
		"Wrap",
		"",
		"func pattern(s *alfa.Server) {\n\talfa.Wrap(s)\n}",
		"func replacement(s *alfa.Server) {\n\talfa.Keep(s)\n}",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.Integer(t, 3, apply.Matched)
	assert.Integer(t, 1, apply.Rewritten)
	refused := 0

	for _, g := range apply.Refused {
		refused += len(g.Locations)
	}

	assert.Integer(t, 2, refused)
	holder := service_tester.ReadFixtureFile(t, d, "pkg/bravo/holder.go")
	assert.StringContains(t, "return alfa.Keep(s)\n", holder)
	assert.StringContains(t, "return alfa.Wrap(alfa.Wrap(s))\n", holder)
	assertCompiles(t, d)
}

func TestApplyPatternCallNeedsExpressionReplacement(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("pattern-expression/src"),
	)
	r, _, e := service_tester.NewService(t).Pattern.ApplyPattern(
		t.Context(),
		d,
		"example/pkg/alfa",
		"New",
		"",
		"func pattern(n string) {\n\talfa.New(n)\n}",
		"func replacement(n string) {\n\t_ = alfa.Named(n)\n}",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, "replacement must be an expression")
}
