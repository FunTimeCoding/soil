package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/naming"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestNameParametersNamesEveryBlankContext(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("name-parameters/src"),
	)
	r, e := o.Symbol.NameParameters(
		t.Context(),
		o.Directory,
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
	source := o.Read("pkg/handler/handler.go")
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
	rewrite_tester.Formatted(t, source)
	assert.StringContains(
		t,
		"func Sub(_ context.Context) {}",
		o.Read("pkg/handler/sub/sub.go"),
	)
}

func TestNameParametersCoversSubtree(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("name-parameters/src"),
	)
	r, e := o.Symbol.NameParameters(
		t.Context(),
		o.Directory,
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
		o.Read("pkg/handler/sub/sub.go"),
	)
}

func TestNameParametersNamesOnlyListedFunctions(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("name-parameters/src"),
	)
	r, e := o.Symbol.NameParameters(
		t.Context(),
		o.Directory,
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
	source := o.Read("pkg/handler/handler.go")
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
	o := service_tester.New(
		t,
		service_tester.WithFixture("name-parameters/src"),
	)
	before := o.Read("pkg/handler/handler.go")
	r, e := o.Symbol.NameParameters(
		t.Context(),
		o.Directory,
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
	assert.StringContains(t, "Get: named x", rewrite_tester.JoinedConcerns(r))
	assert.String(t, before, o.Read("pkg/handler/handler.go"))
}

func TestNameParametersRefusesNameBodyDeclares(t *testing.T) {
	rewrite_tester.NamingRefused(t, "already uses x", "Declared")
}

func TestNameParametersRefusesShadowingPackageName(t *testing.T) {
	rewrite_tester.NamingRefused(t, "already uses x", "Outer")
}

func TestNameParametersRefusesResultOfThatName(t *testing.T) {
	rewrite_tester.NamingRefused(t, "already uses x", "Result")
}

func TestNameParametersRefusesAmbiguousSignature(t *testing.T) {
	rewrite_tester.NamingRefused(t, "ambiguous", "Twice")
}
