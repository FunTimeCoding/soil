package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol_rename"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestRenameSymbols(t *testing.T) {
	d := renameFunctionFixture(t)
	s := service_tester.NewService(t)
	r, e := s.RenameSymbols(
		t.Context(),
		d,
		"example/pkg/target",
		[]*symbol_rename.Rename{
			symbol_rename.New("IsGeneratedHeader", "IsGenerated", ""),
			symbol_rename.New("Check", "Verify", ""),
		},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assertCompiles(t, d)
	declaration := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/is_generated.go",
	)
	assert.StringContains(t, "func IsGenerated(", declaration)
	verify := service_tester.ReadFixtureFile(t, d, "pkg/target/verify.go")
	assert.StringContains(t, "func Verify(", verify)
	assert.StringContains(t, "return IsGenerated(content)", verify)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go")
	assert.StringContains(t, "target.IsGenerated(content)", caller)
}

func TestRenameSymbolsRefusedItemBlocksBatch(t *testing.T) {
	d := renameFunctionFixture(t)
	s := service_tester.NewService(t)
	r, e := s.RenameSymbols(
		t.Context(),
		d,
		"example/pkg/target",
		[]*symbol_rename.Rename{
			symbol_rename.New("IsGeneratedHeader", "IsGenerated", ""),
			symbol_rename.New("Missing", "Something", ""),
		},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "not found")
	declaration := service_tester.ReadFixtureFile(
		t,
		d,
		"pkg/target/is_generated_header.go",
	)
	assert.StringContains(t, "func IsGeneratedHeader(", declaration)
}

func TestRenameSymbolsSameTarget(t *testing.T) {
	d := renameFunctionFixture(t)
	s := service_tester.NewService(t)
	r, e := s.RenameSymbols(
		t.Context(),
		d,
		"example/pkg/target",
		[]*symbol_rename.Rename{
			symbol_rename.New("IsGeneratedHeader", "Verify", ""),
			symbol_rename.New("Check", "Verify", ""),
		},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(
		t,
		r,
		"IsGeneratedHeader and Check both rename to Verify",
	)
}

func TestRenameSymbolsListedTwice(t *testing.T) {
	d := renameFunctionFixture(t)
	s := service_tester.NewService(t)
	r, e := s.RenameSymbols(
		t.Context(),
		d,
		"example/pkg/target",
		[]*symbol_rename.Rename{
			symbol_rename.New("Check", "Verify", ""),
			symbol_rename.New("Check", "Inspect", ""),
		},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "Check is renamed twice")
}

func TestRenameSymbolsOntoFreedName(t *testing.T) {
	d := renameFunctionFixture(t)
	s := service_tester.NewService(t)
	r, e := s.RenameSymbols(
		t.Context(),
		d,
		"example/pkg/target",
		[]*symbol_rename.Rename{
			symbol_rename.New("IsGeneratedHeader", "Check", ""),
			symbol_rename.New("Check", "Verify", ""),
		},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}
