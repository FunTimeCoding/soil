package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"strings"
	"testing"
)

func TestMoveCommentTargetDrift(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-comment/src"))
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		nil,
		"pkg/target/constant.go",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := o.Read("pkg/target/constant/constant.go")
	rewrite_tester.New(t).Formatted(moved)
	assert.StringContains(t, "const Host = \"bravo\" // host trailing", moved)
	assert.StringContains(t, "// Alfa document.", moved)
	assert.StringContains(t, "const Alfa = \"kilo\" // alfa trailing", moved)
	assert.True(
		t,
		strings.Index(moved, "// host trailing") <
			strings.Index(moved, "// Alfa document."),
	)
}

func TestMoveCommentSourceScrub(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-comment/src"))
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/keeper",
		[]string{"moved"},
		"",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := o.Read("pkg/keeper/constant.go")
	rewrite_tester.New(t).Formatted(source)
	assert.StringContains(t, "// Kept document.", source)
	assert.StringContains(t, "const kept = \"golf\" // kept trailing", source)
	assert.False(t, strings.Contains(source, "// Moved document."))
	assert.False(t, strings.Contains(source, "// moved trailing"))
	moved := o.Read("pkg/target/constant/constant.go")
	rewrite_tester.New(t).Formatted(moved)
	assert.StringContains(t, "// Moved document.", moved)
	assert.StringContains(t, "const Moved = \"hotel\" // moved trailing", moved)
}

func TestMoveCommentGroupScrub(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-comment/src"))
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/keeper",
		[]string{"delta"},
		"",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := o.Read("pkg/keeper/constant.go")
	rewrite_tester.New(t).Formatted(source)
	assert.StringContains(t, "echo = \"juliet\"", source)
	assert.StringContains(t, "// Keeper vocabulary.", source)
	assert.False(t, strings.Contains(source, "// Delta document."))
	assert.False(t, strings.Contains(source, "// delta trailing"))
	moved := o.Read("pkg/target/constant/constant.go")
	rewrite_tester.New(t).Formatted(moved)
	assert.StringContains(t, "// Delta document.", moved)
	assert.StringContains(t, "const Delta = \"india\" // delta trailing", moved)
	assert.False(t, strings.Contains(moved, "// Keeper vocabulary."))
}

func TestMoveCommentGroupDocument(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-comment/src"))
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/keeper",
		[]string{"delta", "echo"},
		"",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := o.Read("pkg/keeper/constant.go")
	rewrite_tester.New(t).Formatted(source)
	assert.False(t, strings.Contains(source, "// Keeper vocabulary."))
	moved := o.Read("pkg/target/constant/constant.go")
	rewrite_tester.New(t).Formatted(moved)
	assert.StringContains(t, "// Keeper vocabulary.", moved)
	assert.StringContains(t, "// Delta document.", moved)
	assert.StringContains(t, "Delta = \"india\" // delta trailing", moved)
	assert.StringContains(t, "Echo = \"juliet\"", moved)
}
