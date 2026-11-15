package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/module_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/addition"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestAddParametersPassesArgumentAtEveryCall(t *testing.T) {
	d := service_tester.Prepare(t, "add-parameters/src")
	r, e := service_tester.NewService(t).Symbol.AddParameters(
		t.Context(),
		d,
		[]*addition.Parameter{
			module_tester.ContextParameter("Answer", "", 0, "x"),
			module_tester.ContextParameter("Grouped", "", 1, "x"),
		},
		nil,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	answer := service_tester.ReadFixtureFile(t, d, "pkg/target/answer.go")
	assert.StringContains(t, "import \"context\"", answer)
	assert.StringContains(
		t,
		"func Answer(\n\tx context.Context,\n\tmessage string,\n\tcode int,\n) string {",
		answer,
	)
	rewrite_tester.Formatted(t, answer)
	target := service_tester.ReadFixtureFile(t, d, "pkg/target/target.go")
	assert.StringContains(
		t,
		"func Grouped(first int, x context.Context, second int) int {",
		target,
	)
	rewrite_tester.Formatted(t, target)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/caller.go")
	assert.StringContains(t, `target.Answer(x, "a", 1)`, caller)
	assert.StringContains(
		t,
		"target.Answer(\n\t\tx,\n\t\t\"b\",\n\t\t2,\n\t)",
		caller,
	)
	assert.StringContains(t, "target.Grouped(1, x, 2)", caller)
	rewrite_tester.Formatted(t, caller)
}

func TestAddParametersImportsPackageCallSiteLacks(t *testing.T) {
	d := service_tester.Prepare(t, "add-parameters/src")
	r, e := service_tester.NewService(t).Symbol.AddParameters(
		t.Context(),
		d,
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
	other := service_tester.ReadFixtureFile(t, d, "pkg/other/other.go")
	assert.StringContains(t, "\"context\"", other)
	assert.StringContains(t, `target.Plain(context.Background(), "m")`, other)
	rewrite_tester.Formatted(t, other)
}

func TestAddParametersDryRunWritesNothing(t *testing.T) {
	d := service_tester.Prepare(t, "add-parameters/src")
	before := service_tester.ReadFixtureFile(t, d, "pkg/caller/caller.go")
	r, e := service_tester.NewService(t).Symbol.AddParameters(
		t.Context(),
		d,
		[]*addition.Parameter{
			module_tester.ContextParameter("Answer", "", 0, "x"),
		},
		nil,
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(t, "passed x", rewrite_tester.JoinedConcerns(r))
	assert.String(
		t,
		before,
		service_tester.ReadFixtureFile(t, d, "pkg/caller/caller.go"),
	)
}

func TestAddParametersRewritesReplacingModule(t *testing.T) {
	for _, full := range []bool{false, true} {
		library, user := module_tester.ReplacedLibrary(
			t,
			"add-parameters-replacer",
		)
		s := module_tester.ReplacingService(t, library, user)

		if full {
			s.Workspace.UseFullLoad()
		}

		r, e := s.Symbol.AddParameters(
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
			service_tester.ReadFixtureFile(t, library, "lib.go"),
		)
		assert.StringContains(
			t,
			`lib.Log(x, "a")`,
			service_tester.ReadFixtureFile(t, user, "pkg/user/user.go"),
		)
	}
}

func TestAddParametersRefusesArgumentOutOfScope(t *testing.T) {
	rewrite_tester.AdditionRefused(
		t,
		"does not resolve here",
		module_tester.ContextParameter("Unbound", "", 0, "x"),
	)
}

func TestAddParametersRefusesArgumentOfWrongType(t *testing.T) {
	rewrite_tester.AdditionRefused(
		t,
		"is not assignable",
		module_tester.ContextParameter("Mistyped", "", 0, "x"),
	)
}

func TestAddParametersRefusesWiderArgumentThanParameter(t *testing.T) {
	rewrite_tester.AdditionRefused(
		t,
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
	rewrite_tester.AdditionRefused(
		t,
		"used as a value",
		module_tester.ContextParameter("Valued", "", 0, "x"),
	)
}

func TestAddParametersRefusesInterfaceMethod(t *testing.T) {
	rewrite_tester.AdditionRefused(
		t,
		"satisfies",
		module_tester.ContextParameter("Reply", "Server", 0, "x"),
	)
}

func TestAddParametersRefusesNameBodyUses(t *testing.T) {
	rewrite_tester.AdditionRefused(
		t,
		"already uses x",
		module_tester.ContextParameter("Shadow", "", 0, "x"),
	)
}

func TestAddParametersRefusesPositionAfterVariadic(t *testing.T) {
	rewrite_tester.AdditionRefused(
		t,
		"follows the variadic",
		module_tester.ContextParameter("Log", "", 2, "x"),
	)
}

func TestAddParametersRefusesPositionOutOfRange(t *testing.T) {
	rewrite_tester.AdditionRefused(
		t,
		"out of range",
		module_tester.ContextParameter("Plain", "", 3, "x"),
	)
}

func TestAddParametersRefusesUnnamedParameters(t *testing.T) {
	rewrite_tester.AdditionRefused(
		t,
		"unnamed parameters",
		module_tester.ContextParameter("Unnamed", "", 0, "x"),
	)
}
