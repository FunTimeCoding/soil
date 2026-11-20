package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/module_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/addition"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestAddParametersPassesArgumentAtEveryCall(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("add-parameters/src"))
	r, e := o.Symbol.AddParameters(
		t.Context(),
		o.Directory,
		[]*addition.Parameter{
			module_tester.ContextParameter("Answer", "", 0, "x"),
			module_tester.ContextParameter("Grouped", "", 1, "x"),
		},
		nil,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	answer := o.Read("pkg/target/answer.go")
	assert.StringContains(t, "import \"context\"", answer)
	assert.StringContains(
		t,
		"func Answer(\n\tx context.Context,\n\tmessage string,\n\tcode int,\n) string {",
		answer,
	)
	rewrite_tester.New(t).Formatted(answer)
	target := o.Read("pkg/target/target.go")
	assert.StringContains(
		t,
		"func Grouped(first int, x context.Context, second int) int {",
		target,
	)
	rewrite_tester.New(t).Formatted(target)
	caller := o.Read("pkg/caller/caller.go")
	assert.StringContains(t, `target.Answer(x, "a", 1)`, caller)
	assert.StringContains(
		t,
		"target.Answer(\n\t\tx,\n\t\t\"b\",\n\t\t2,\n\t)",
		caller,
	)
	assert.StringContains(t, "target.Grouped(1, x, 2)", caller)
	rewrite_tester.New(t).Formatted(caller)
}

func TestAddParametersImportsPackageCallSiteLacks(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("add-parameters/src"))
	r, e := o.Symbol.AddParameters(
		t.Context(),
		o.Directory,
		[]*addition.Parameter{
			module_tester.ContextParameter(
				"Plain",
				"",
				0,
				"context.Background()",
			),
		},
		nil,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	other := o.Read("pkg/other/other.go")
	assert.StringContains(t, "\"context\"", other)
	assert.StringContains(t, `target.Plain(context.Background(), "m")`, other)
	rewrite_tester.New(t).Formatted(other)
}

func TestAddParametersDryRunWritesNothing(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("add-parameters/src"))
	before := o.Read("pkg/caller/caller.go")
	r, e := o.Symbol.AddParameters(
		t.Context(),
		o.Directory,
		[]*addition.Parameter{
			module_tester.ContextParameter("Answer", "", 0, "x"),
		},
		nil,
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(t, "passed x", rewrite_tester.JoinedConcerns(r))
	assert.String(t, before, o.Read("pkg/caller/caller.go"))
}

func TestAddParametersRewritesReplacingModule(t *testing.T) {
	for _, full := range []bool{false, true} {
		library, o := module_tester.New(t).ReplacedLibrary(
			"add-parameters-replacer",
		)

		if full {
			o.Workspace.UseFullLoad()
		}

		r, e := o.Symbol.AddParameters(
			t.Context(),
			library,
			[]*addition.Parameter{
				addition.NewParameter(
					"other.test/lib",
					"Log",
					"",
					"x",
					"context.Context",
					0,
					"x",
				),
			},
			nil,
			false,
		)
		assert.FatalOnError(t, e)
		testutil.AssertBlocked(t, r, 0)
		assert.String(
			t,
			"package lib\n\nimport \"context\"\n\nfunc Log(\n\tx context.Context,\n\tprefix string,\n) string {\n\treturn prefix\n}\n",
			system_tester.New(t).Read(library, "lib.go"),
		)
		assert.StringContains(t, `lib.Log(x, "a")`, o.Read("pkg/user/user.go"))
	}
}

func TestAddParametersRefusesArgumentOutOfScope(t *testing.T) {
	rewrite_tester.New(t).AdditionRefused(
		"does not resolve here",
		module_tester.ContextParameter("Unbound", "", 0, "x"),
	)
}

func TestAddParametersRefusesArgumentOfWrongType(t *testing.T) {
	rewrite_tester.New(t).AdditionRefused(
		"is not assignable",
		module_tester.ContextParameter("Mistyped", "", 0, "x"),
	)
}

func TestAddParametersRefusesWiderArgumentThanParameter(t *testing.T) {
	rewrite_tester.New(t).AdditionRefused(
		"is not assignable",
		addition.NewParameter(
			"example/pkg/target",
			"Narrowed",
			"",
			"title",
			"string",
			0,
			"value",
		),
	)
}

func TestAddParametersRefusesValueUse(t *testing.T) {
	rewrite_tester.New(t).AdditionRefused(
		"used as a value",
		module_tester.ContextParameter("Valued", "", 0, "x"),
	)
}

func TestAddParametersRefusesInterfaceMethod(t *testing.T) {
	rewrite_tester.New(t).AdditionRefused(
		"satisfies",
		module_tester.ContextParameter("Reply", "Server", 0, "x"),
	)
}

func TestAddParametersRefusesNameBodyUses(t *testing.T) {
	rewrite_tester.New(t).AdditionRefused(
		"already uses x",
		module_tester.ContextParameter("Shadow", "", 0, "x"),
	)
}

func TestAddParametersRefusesPositionAfterVariadic(t *testing.T) {
	rewrite_tester.New(t).AdditionRefused(
		"follows the variadic",
		module_tester.ContextParameter("Log", "", 2, "x"),
	)
}

func TestAddParametersRefusesPositionOutOfRange(t *testing.T) {
	rewrite_tester.New(t).AdditionRefused(
		"out of range",
		module_tester.ContextParameter("Plain", "", 3, "x"),
	)
}

func TestAddParametersRefusesUnnamedParameters(t *testing.T) {
	rewrite_tester.New(t).AdditionRefused(
		"unnamed parameters",
		module_tester.ContextParameter("Unnamed", "", 0, "x"),
	)
}
