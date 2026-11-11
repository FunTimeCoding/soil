package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestMatchPatternComparesQualifiersByPackage(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("pattern-aliased/src"),
	)
	r, match, e := service_tester.NewService(t).Pattern.MatchPattern(
		d,
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
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("pattern-aliased/src"),
	)
	r, _, e := service_tester.NewService(t).Pattern.ApplyPattern(
		t.Context(),
		d,
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
		service_tester.ReadFixtureFile(t, d, "pkg/charlie/run.go"),
	)
	assert.StringContains(
		t,
		"return alfa.Advance(n, deltaConstant.Step)\n",
		service_tester.ReadFixtureFile(t, d, "pkg/foxtrot/run.go"),
	)
	assertCompiles(t, d)
}
