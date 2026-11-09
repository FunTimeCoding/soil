package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/naming"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestNameParametersNamesEveryBlankContext(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("name-parameters/src"),
	)
	r, e := service_tester.NewService(t).Symbol.NameParameters(
		t.Context(),
		d,
		naming.NewRequest(
			"example/pkg/handler",
			nil,
			"context.Context",
			"x",
			nil,
		),
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := service_tester.ReadFixtureFile(t, d, "pkg/handler/handler.go")
	assert.StringContains(
		t,
		"func (s *Server) Get(\n\tx context.Context,\n\tidentifier string,\n) string {",
		source,
	)
	assert.StringContains(
		t,
		"func (s *Server) List(x context.Context) []string {",
		source,
	)
	assert.StringContains(
		t,
		"func Pair(x context.Context, _ string) {}",
		source,
	)
	assert.StringContains(
		t,
		"func Named(y context.Context) context.Context {",
		source,
	)
	assert.StringContains(t, "return func(x context.Context) string {", source)
	assert.StringContains(
		t,
		"func Literal() func(context.Context) string {",
		source,
	)
	assertFormatted(t, source)
	assert.StringContains(
		t,
		"func Sub(_ context.Context) {}",
		service_tester.ReadFixtureFile(t, d, "pkg/handler/sub/sub.go"),
	)
}

func TestNameParametersCoversSubtree(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("name-parameters/src"),
	)
	r, e := service_tester.NewService(t).Symbol.NameParameters(
		t.Context(),
		d,
		naming.NewRequest(
			"example/pkg/handler/...",
			nil,
			"context.Context",
			"x",
			nil,
		),
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(
		t,
		"func Sub(x context.Context) {}",
		service_tester.ReadFixtureFile(t, d, "pkg/handler/sub/sub.go"),
	)
}

func TestNameParametersNamesOnlyListedFunctions(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("name-parameters/src"),
	)
	r, e := service_tester.NewService(t).Symbol.NameParameters(
		t.Context(),
		d,
		naming.NewRequest(
			"example/pkg/handler",
			[]*naming.Function{naming.NewFunction("Get", "Server")},
			"context.Context",
			"x",
			nil,
		),
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := service_tester.ReadFixtureFile(t, d, "pkg/handler/handler.go")
	assert.StringContains(
		t,
		"\tx context.Context,\n\tidentifier string,",
		source,
	)
	assert.StringContains(
		t,
		"func (s *Server) List(context.Context) []string {",
		source,
	)
}

func TestNameParametersDryRunWritesNothing(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("name-parameters/src"),
	)
	before := service_tester.ReadFixtureFile(t, d, "pkg/handler/handler.go")
	r, e := service_tester.NewService(t).Symbol.NameParameters(
		t.Context(),
		d,
		naming.NewRequest(
			"example/pkg/handler",
			nil,
			"context.Context",
			"x",
			nil,
		),
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(t, "Get: named x", joinedConcerns(r))
	assert.String(
		t,
		before,
		service_tester.ReadFixtureFile(t, d, "pkg/handler/handler.go"),
	)
}

func TestNameParametersRefusesNameBodyDeclares(t *testing.T) {
	assertNamingRefused(t, "already uses x", "Declared")
}

func TestNameParametersRefusesShadowingPackageName(t *testing.T) {
	assertNamingRefused(t, "already uses x", "Outer")
}

func TestNameParametersRefusesResultOfThatName(t *testing.T) {
	assertNamingRefused(t, "already uses x", "Result")
}

func TestNameParametersRefusesAmbiguousSignature(t *testing.T) {
	assertNamingRefused(t, "ambiguous", "Twice")
}
