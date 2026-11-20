package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBatchFile(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("batch-file/src"))
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
	_, e = os.Stat(filepath.Join(o.Directory, "pkg/target/constant.go"))
	assert.True(t, os.IsNotExist(e))
	moved := o.Read("pkg/target/constant/constant.go")
	rewrite_tester.New(t).Formatted(moved)
	assert.Integer(t, 1, strings.Count(moved, "const ("))
	assert.StringContains(t, "First  = \"alfa\"", moved)
	assert.StringContains(t, "Second = \"bravo\"", moved)
	assert.StringContains(t, "Third  = \"charlie\"", moved)
	run := o.Read("pkg/target/run.go")
	rewrite_tester.New(t).Formatted(run)
	assert.StringContains(
		t,
		"constant.First + constant.Second + constant.Third",
		run,
	)
}

func TestBatchEnum(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("batch-enum/src"))
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
	assert.StringContains(t, "type Mode int", moved)
	assert.StringContains(t, "Off Mode = iota", moved)
	assert.False(t, strings.Contains(moved, "constant.Mode"))
	run := o.Read("pkg/target/run.go")
	rewrite_tester.New(t).Formatted(run)
	assert.StringContains(t, "func Run() constant.Mode", run)
	assert.StringContains(t, "return constant.On", run)
	caller := o.Read("pkg/caller/run.go")
	rewrite_tester.New(t).Formatted(caller)
	assert.StringContains(t, "func Run() targetConstant.Mode", caller)
	assert.StringContains(t, "targetConstant.Off == 0", caller)
	assert.False(t, strings.Contains(caller, "target."))
}

func TestBatchSubset(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("batch-file/src"))
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		[]string{"first", "second"},
		"",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := o.Read("pkg/target/constant.go")
	rewrite_tester.New(t).Formatted(source)
	assert.StringContains(t, "third = \"charlie\"", source)
	assert.False(t, strings.Contains(source, "first"))
	moved := o.Read("pkg/target/constant/constant.go")
	rewrite_tester.New(t).Formatted(moved)
	assert.StringContains(t, "First  = \"alfa\"", moved)
	assert.StringContains(t, "Second = \"bravo\"", moved)
	run := o.Read("pkg/target/run.go")
	assert.StringContains(t, "constant.First + constant.Second + third", run)
}

func TestBatchCollision(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("batch-collision/src"),
	)
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
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "both move as Mode")
}

func TestBatchDependencyTogether(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("move-dependency/src"),
	)
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		[]string{"prefix", "composed"},
		"",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := o.Read("pkg/target/constant/constant.go")
	rewrite_tester.New(t).Formatted(moved)
	assert.StringContains(t, "composed = prefix + \"bravo\"", moved)
}

func TestBatchMultiName(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("batch-multi-name/src"),
	)
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
	assert.StringContains(t, "Alfa, Bravo = \"a\", \"b\"", moved)
	run := o.Read("pkg/target/run.go")
	assert.StringContains(t, "constant.Alfa + constant.Bravo", run)
}

func TestBatchMultiNamePartial(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("batch-multi-name/src"),
	)
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		[]string{"alfa"},
		"",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "move them together")
}

func TestBatchMethodRefused(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("batch-method/src"))
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		nil,
		"pkg/target/store.go",
		"example/pkg/target/constant",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(
		t,
		r,
		"move it with move_methods, or its type with extract_type",
	)
}

func TestBatchMethodSetRefused(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("batch-method/src"))
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		[]string{"Store"},
		"",
		"example/pkg/target/constant",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "Store has methods")
	testutil.AssertBlockedContains(t, r, constant.ExtractType)
}

func TestBackReferenceQualified(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("back-reference/src"))
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/source",
		[]string{"Mover"},
		"",
		"example/pkg/target",
		"mover.go",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := o.Read("pkg/target/mover.go")
	rewrite_tester.New(t).Formatted(moved)
	assert.StringContains(t, "example/pkg/source", moved)
	assert.StringContains(t, "source.Stayer()", moved)
	source := o.Read("pkg/source/source.go")
	rewrite_tester.New(t).Formatted(source)
	assert.StringContains(t, "func Stayer()", source)
	assert.False(t, strings.Contains(source, "func Mover"))
}

func TestBackReferenceCycleRefused(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("cycle/src"))
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/source",
		[]string{"Mover"},
		"",
		"example/pkg/target",
		"mover.go",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "cycle")
}

func TestBackReferenceUnexportedIsExported(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("back-reference-unexported/src"),
	)
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/source",
		[]string{"Mover"},
		"",
		"example/pkg/target",
		"mover.go",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	rewrite_tester.New(t).FixedContains("stayer → Stayer", r)
	rewrite_tester.New(t).FixedContains(
		"stayer → source.Stayer (back reference)",
		r,
	)
	assert.String(
		t,
		"package source\n\nfunc Stayer() string {\n\treturn \"stay\"\n}\n",
		o.Read("pkg/source/source.go"),
	)
	assert.StringContains(
		t,
		"return source.Stayer() + \" moved\"",
		o.Read("pkg/target/mover.go"),
	)
	rewrite_tester.New(t).Compiles(o.Directory)
}

func TestBackReferenceExportCollisionRefused(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("back-reference-collision/src"),
	)
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/source",
		[]string{"Mover"},
		"",
		"example/pkg/target",
		"mover.go",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "Stayer already exists")
}
