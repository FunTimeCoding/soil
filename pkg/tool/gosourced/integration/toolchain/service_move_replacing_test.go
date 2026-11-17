package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/module_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"testing"
)

func TestMoveSymbolsRequalifiesReplacingModule(t *testing.T) {
	library, o := module_tester.ReplacedLibrary(t, "move-replacer")
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		library,
		"other.test/lib/source",
		[]string{"Helper"},
		"",
		"other.test/lib/target",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"package user\n\nimport \"other.test/lib/target\"\n\nfunc Use() string {\n\treturn target.Helper()\n}\n",
		o.Read("pkg/user/user.go"),
	)
	both := o.Read("pkg/both/both.go")
	rewrite_tester.Formatted(t, both)
	assert.StringContains(t, "\"other.test/lib/source\"", both)
	assert.StringContains(t, "\"other.test/lib/target\"", both)
	assert.StringContains(t, "return target.Helper() + source.Keep()", both)
	clash := o.Read("pkg/clash/clash.go")
	rewrite_tester.Formatted(t, clash)
	assert.StringContains(t, "lib \"other.test/lib/target\"", clash)
	assert.StringContains(t, "return target.Local() + lib.Helper()", clash)
}

func TestMoveSymbolsToNewPackageRequalifiesReplacingModule(t *testing.T) {
	library, o := module_tester.ReplacedLibrary(t, "move-replacer")
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		library,
		"other.test/lib/source",
		[]string{"Helper"},
		"",
		"other.test/lib/fresh",
		"",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"package user\n\nimport \"other.test/lib/fresh\"\n\nfunc Use() string {\n\treturn fresh.Helper()\n}\n",
		o.Read("pkg/user/user.go"),
	)
}

func TestMoveSymbolsByFileRequalifiesReplacingModule(t *testing.T) {
	library, o := module_tester.ReplacedLibrary(t, "move-replacer")
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		library,
		"other.test/lib/source",
		nil,
		"source/source.go",
		"other.test/lib/target",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"package both\n\nimport \"other.test/lib/target\"\n\nfunc Both() string {\n\treturn target.Helper() + target.Keep()\n}\n",
		o.Read("pkg/both/both.go"),
	)
	widget := o.Read("pkg/widget/widget.go")
	rewrite_tester.Formatted(t, widget)
	assert.StringContains(t, "return w.Run() + target.Limit", widget)
	assert.StringContains(t, "\tsource.Widget\n", widget)
}

func TestExtractTypeRequalifiesReplacingModule(t *testing.T) {
	library, o := module_tester.ReplacedLibrary(t, "move-replacer")
	r, e := o.Symbol.ExtractType(
		t.Context(),
		library,
		"other.test/lib/source",
		"Widget",
		"other.test/lib/target",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	widget := o.Read("pkg/widget/widget.go")
	rewrite_tester.Formatted(t, widget)
	assert.StringContains(t, "\ttarget.Widget\n", widget)
	assert.StringContains(t, "w := &target.Widget{Size: 1}", widget)
	assert.StringContains(t, "return w.Run() + source.Limit", widget)
}

func TestMoveSymbolsRequalifiesReplacingModuleOnWholeLoad(t *testing.T) {
	library, o := module_tester.ReplacedLibrary(t, "move-replacer")
	o.Workspace.UseFullLoad()
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		library,
		"other.test/lib/source",
		nil,
		"source/source.go",
		"other.test/lib/target",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(
		t,
		"return target.Helper() + target.Keep()",
		o.Read("pkg/both/both.go"),
	)
}

func TestMoveSymbolsDryRunLeavesReplacingModuleUntouched(t *testing.T) {
	library, o := module_tester.ReplacedLibrary(t, "move-replacer")
	before := rewrite_tester.DirectoryTree(t, o.Directory)
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		library,
		"other.test/lib/source",
		[]string{"Helper"},
		"",
		"other.test/lib/target",
		"",
		false,
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(t, before, rewrite_tester.DirectoryTree(t, o.Directory))
	assert.StringContains(
		t,
		"pkg/user/user.go: Helper → target.Helper",
		rewrite_tester.JoinedConcerns(r),
	)
}
