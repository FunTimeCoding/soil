package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/addition"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestAddParametersPassesTheArgumentAtEveryCall(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("add-parameters/src"),
	)
	r, e := service_tester.NewService(t).Symbol.AddParameters(
		t.Context(),
		d,
		[]*addition.Parameter{
			contextParameter("Answer", "", 0, "x"),
			contextParameter("Grouped", "", 1, "x"),
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
	assertFormatted(t, answer)
	target := service_tester.ReadFixtureFile(t, d, "pkg/target/target.go")
	assert.StringContains(
		t,
		"func Grouped(first int, x context.Context, second int) int {",
		target,
	)
	assertFormatted(t, target)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/caller.go")
	assert.StringContains(t, `target.Answer(x, "a", 1)`, caller)
	assert.StringContains(
		t,
		"target.Answer(\n\t\tx,\n\t\t\"b\",\n\t\t2,\n\t)",
		caller,
	)
	assert.StringContains(t, "target.Grouped(1, x, 2)", caller)
	assertFormatted(t, caller)
}

func TestAddParametersImportsAPackageTheCallSiteLacks(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("add-parameters/src"),
	)
	r, e := service_tester.NewService(t).Symbol.AddParameters(
		t.Context(),
		d,
		[]*addition.Parameter{
			contextParameter("Plain", "", 0, "context.Background()"),
		},
		nil,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	other := service_tester.ReadFixtureFile(t, d, "pkg/other/other.go")
	assert.StringContains(t, "\"context\"", other)
	assert.StringContains(t, `target.Plain(context.Background(), "m")`, other)
	assertFormatted(t, other)
}

func TestAddParametersDryRunWritesNothing(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("add-parameters/src"),
	)
	before := service_tester.ReadFixtureFile(t, d, "pkg/caller/caller.go")
	r, e := service_tester.NewService(t).Symbol.AddParameters(
		t.Context(),
		d,
		[]*addition.Parameter{contextParameter("Answer", "", 0, "x")},
		nil,
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(t, "passed x", joinedConcerns(r))
	assert.String(
		t,
		before,
		service_tester.ReadFixtureFile(t, d, "pkg/caller/caller.go"),
	)
}

func TestAddParametersRewritesAReplacingModule(t *testing.T) {
	for _, full := range []bool{false, true} {
		library, user := prepareReplacedLibrary(t, "add-parameters-replacer")
		s := replacingService(t, library, user)

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

func TestAddParametersRefusesAnArgumentOutOfScope(t *testing.T) {
	assertAdditionRefused(
		t,
		"does not resolve here",
		contextParameter("Unbound", "", 0, "x"),
	)
}

func TestAddParametersRefusesAnArgumentOfTheWrongType(t *testing.T) {
	assertAdditionRefused(
		t,
		"is not assignable",
		contextParameter("Mistyped", "", 0, "x"),
	)
}

func TestAddParametersRefusesAWiderArgumentThanTheParameter(t *testing.T) {
	assertAdditionRefused(
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

func TestAddParametersRefusesAValueUse(t *testing.T) {
	assertAdditionRefused(
		t,
		"used as a value",
		contextParameter("Valued", "", 0, "x"),
	)
}

func TestAddParametersRefusesAnInterfaceMethod(t *testing.T) {
	assertAdditionRefused(
		t,
		"satisfies",
		contextParameter("Reply", "Server", 0, "x"),
	)
}

func TestAddParametersRefusesANameTheBodyUses(t *testing.T) {
	assertAdditionRefused(
		t,
		"already uses x",
		contextParameter("Shadow", "", 0, "x"),
	)
}

func TestAddParametersRefusesAPositionAfterTheVariadic(t *testing.T) {
	assertAdditionRefused(
		t,
		"follows the variadic",
		contextParameter("Log", "", 2, "x"),
	)
}

func TestAddParametersRefusesAPositionOutOfRange(t *testing.T) {
	assertAdditionRefused(
		t,
		"out of range",
		contextParameter("Plain", "", 3, "x"),
	)
}

func TestAddParametersRefusesUnnamedParameters(t *testing.T) {
	assertAdditionRefused(
		t,
		"unnamed parameters",
		contextParameter("Unnamed", "", 0, "x"),
	)
}
