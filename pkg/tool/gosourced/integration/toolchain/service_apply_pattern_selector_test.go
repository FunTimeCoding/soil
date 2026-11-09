package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestApplyPatternRewritesSelectorChainBeforeAnchor(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("selector/src"),
	)
	r, apply, e := service_tester.NewService(t).Pattern.ApplyPattern(
		t.Context(),
		d,
		"example/pkg/alfa",
		"Run",
		"Alfa",
		"func pattern(b *bravo.Bravo, a []any) {\n\tx := b.Run(a...)\n}",
		"func replacement(b *bravo.Bravo, a []any) {\n\tx := b.Named.Run(a...)\n}",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.Integer(t, 1, apply.Rewritten)
	assert.True(t, apply.Applied)
	assert.String(
		t,
		"package charlie\n\nimport \"example/pkg/bravo\"\n\nfunc Charlie(b *bravo.Bravo) int {\n\tx := b.Named.Run(1, 2)\n\n\treturn x\n}\n",
		service_tester.ReadFixtureFile(t, d, "pkg/charlie/charlie.go"),
	)
}

func TestApplyPatternClaimsNoRewriteThatChangesNothing(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("selector/src"),
	)
	r, apply, e := service_tester.NewService(t).Pattern.ApplyPattern(
		t.Context(),
		d,
		"example/pkg/alfa",
		"Run",
		"Alfa",
		"func pattern(b *bravo.Bravo, a []any) {\n\tx := b.Run(a...)\n}",
		"func replacement(b *bravo.Bravo, a []any) {\n\tx := b.Run(a...)\n}",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.Integer(t, 1, apply.Matched)
	assert.Integer(t, 0, apply.Rewritten)
	assert.False(t, apply.Applied)
	assert.StringContains(t, "changes no file", apply.Refusal)
}
