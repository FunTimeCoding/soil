package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestMoveKeepsVariadicSpread(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-variadic/src"),
	)
	s := service_tester.NewService(t)
	r, e := s.Symbol.MoveSymbols(
		t.Context(),
		d,
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
	user := service_tester.ReadFixtureFile(t, d, "pkg/user/use.go")
	assertFormatted(t, user)
	assert.StringContains(
		t,
		"destination.AdditionalEntriesForTheVariadicSpread...)",
		user,
	)
}
