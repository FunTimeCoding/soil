package string_constant

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/string_constant"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestFlagged(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/flagged")
	string_constant.Check(p, results)
	testutil.AssertBlocked(t, results, 4)
	testutil.AssertBlockedContains(
		t,
		results,
		`string literal "name" has constants constant.Name or constant.NameColumn`,
	)
	testutil.AssertBlockedContains(
		t,
		results,
		`string literal "query" has constant constant.Query`,
	)
}

func TestClean(t *testing.T) {
	temporary := testutil.PrepareTestPackage(t, "testdata/src/clean")
	testutil.WriteModFile(t, temporary, "clean")
	p, results := testutil.LoadFromDirectory(t, temporary)
	string_constant.Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}

func TestNoConstant(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/no_constant")
	string_constant.Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}

func TestExpected(t *testing.T) {
	temporary := testutil.PrepareTestPackage(t, "testdata/src/expected")
	testutil.WriteModFile(t, temporary, "expected.test")
	p, results := testutil.LoadFromDirectory(t, temporary)
	string_constant.Check(p, results)
	testutil.AssertBlocked(t, results, 5)
}

func TestTester(t *testing.T) {
	temporary := testutil.PrepareTestPackage(t, "testdata/src/tester")
	testutil.WriteModFile(t, temporary, "tester.test")
	p, results := testutil.LoadFromDirectory(t, temporary)
	string_constant.Check(p, results)
	testutil.AssertBlocked(t, results, 2)
	testutil.AssertBlockedContains(
		t,
		results,
		`constant constant.Name in expected value should be a literal`,
	)
	testutil.AssertBlockedContains(
		t,
		results,
		`string literal "name" has constant constant.Name`,
	)
}
