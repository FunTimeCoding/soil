package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestMoveSymbolsRequalifiesAReplacingModule(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "move-replacer")
	r, e := replacingService(library, user).MoveSymbols(
		library,
		"other.test/lib/source",
		[]string{"Helper"},
		"",
		"other.test/lib/target",
		"",
		false,
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"package user\n\nimport \"other.test/lib/target\"\n\nfunc Use() string {\n\treturn target.Helper()\n}\n",
		service_tester.ReadFixtureFile(t, user, "pkg/user/user.go"),
	)
	both := service_tester.ReadFixtureFile(t, user, "pkg/both/both.go")
	assertFormatted(t, both)
	assert.StringContains(t, "\"other.test/lib/source\"", both)
	assert.StringContains(t, "\"other.test/lib/target\"", both)
	assert.StringContains(t, "return target.Helper() + source.Keep()", both)
	clash := service_tester.ReadFixtureFile(t, user, "pkg/clash/clash.go")
	assertFormatted(t, clash)
	assert.StringContains(t, "lib \"other.test/lib/target\"", clash)
	assert.StringContains(t, "return target.Local() + lib.Helper()", clash)
}

func TestMoveSymbolsToANewPackageRequalifiesAReplacingModule(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "move-replacer")
	r, e := replacingService(library, user).MoveSymbols(
		library,
		"other.test/lib/source",
		[]string{"Helper"},
		"",
		"other.test/lib/fresh",
		"",
		true,
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"package user\n\nimport \"other.test/lib/fresh\"\n\nfunc Use() string {\n\treturn fresh.Helper()\n}\n",
		service_tester.ReadFixtureFile(t, user, "pkg/user/user.go"),
	)
}

func TestMoveSymbolsByFileRequalifiesAReplacingModule(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "move-replacer")
	r, e := replacingService(library, user).MoveSymbols(
		library,
		"other.test/lib/source",
		nil,
		"source/source.go",
		"other.test/lib/target",
		"",
		false,
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"package both\n\nimport \"other.test/lib/target\"\n\nfunc Both() string {\n\treturn target.Helper() + target.Keep()\n}\n",
		service_tester.ReadFixtureFile(t, user, "pkg/both/both.go"),
	)
	widget := service_tester.ReadFixtureFile(t, user, "pkg/widget/widget.go")
	assertFormatted(t, widget)
	assert.StringContains(t, "return w.Run() + target.Limit", widget)
	assert.StringContains(t, "\tsource.Widget\n", widget)
}

func TestExtractTypeRequalifiesAReplacingModule(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "move-replacer")
	r, e := replacingService(library, user).ExtractType(
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
	widget := service_tester.ReadFixtureFile(t, user, "pkg/widget/widget.go")
	assertFormatted(t, widget)
	assert.StringContains(t, "\ttarget.Widget\n", widget)
	assert.StringContains(t, "w := &target.Widget{Size: 1}", widget)
	assert.StringContains(t, "return w.Run() + source.Limit", widget)
}

func TestMoveSymbolsRequalifiesAReplacingModuleOnTheWholeLoad(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "move-replacer")
	s := replacingService(library, user)
	s.UseFullLoad()
	r, e := s.MoveSymbols(
		library,
		"other.test/lib/source",
		nil,
		"source/source.go",
		"other.test/lib/target",
		"",
		false,
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(
		t,
		"return target.Helper() + target.Keep()",
		service_tester.ReadFixtureFile(t, user, "pkg/both/both.go"),
	)
}

func TestMoveSymbolsDryRunLeavesAReplacingModuleUntouched(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "move-replacer")
	before := treeOf(t, user)
	r, e := replacingService(library, user).MoveSymbols(
		library,
		"other.test/lib/source",
		[]string{"Helper"},
		"",
		"other.test/lib/target",
		"",
		false,
		false,
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(t, before, treeOf(t, user))
	assert.StringContains(
		t,
		"pkg/user/user.go: Helper → target.Helper",
		joinedConcerns(r),
	)
}
