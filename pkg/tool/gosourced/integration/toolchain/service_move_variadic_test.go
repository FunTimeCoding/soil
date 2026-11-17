package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestMoveKeepsVariadicSpread(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("move-variadic/src"))
	r, e := o.Symbol.MoveSymbols(
		t.Context(),
		o.Directory,
		"example/pkg/source",
		[]string{"AdditionalEntriesForTheVariadicSpread"},
		"",
		"example/pkg/destination",
		"list.go",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	user := o.Read("pkg/user/use.go")
	rewrite_tester.Formatted(t, user)
	assert.StringContains(
		t,
		"destination.AdditionalEntriesForTheVariadicSpread...)",
		user,
	)
}
