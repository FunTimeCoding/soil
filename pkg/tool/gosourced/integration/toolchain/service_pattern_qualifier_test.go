package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestMatchPatternComparesQualifiersByPackage(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("pattern-aliased/src"),
	)
	r, match, e := o.Pattern.MatchPattern(
		o.Directory,
		"example/pkg/alfa",
		"Advance",
		"",
		"import deltaConstant \"example/pkg/delta/constant\"\n\nfunc pattern(n int) int {\n\treturn alfa.Advance(n, deltaConstant.Step)\n}",
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.Integer(t, 3, match.Total)
	assert.Integer(t, 2, match.Matched)
	assert.Integer(t, 1, len(match.Unmatched))
	assert.String(t, "pkg/foxtrot/run.go", match.Unmatched[0].Locations[0].File)
}

func TestApplyPatternKeepsSiteAliases(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("pattern-aliased/src"),
	)
	r, _, e := o.Pattern.ApplyPattern(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"Advance",
		"",
		"import deltaConstant \"example/pkg/delta/constant\"\n\nfunc pattern(n int) int {\n\treturn alfa.Advance(n, deltaConstant.Step)\n}",
		"import deltaConstant \"example/pkg/delta/constant\"\n\nfunc replacement(n int) int {\n\treturn alfa.Advance(n, deltaConstant.Step) * 2\n}",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(
		t,
		"return stepper.Advance(n, stepConstant.Step) * 2",
		o.Read("pkg/charlie/run.go"),
	)
	assert.StringContains(
		t,
		"return alfa.Advance(n, deltaConstant.Step)\n",
		o.Read("pkg/foxtrot/run.go"),
	)
	rewrite_tester.Compiles(t, o.Directory)
}
