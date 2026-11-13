package string_constant

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/string_constant"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"path/filepath"
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

func TestFlaggedFromOwnConstantFile(t *testing.T) {
	directory := t.TempDir()
	testutil.WriteModFile(t, directory, "alfa")
	testutil.WriteFile(
		t,
		directory,
		"constant.go",
		"package alfa\n\nconst Bravo = \"charlie delta\"\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"echo.go",
		"package alfa\n\nfunc Echo() {\n\tprintln(\"charlie delta\")\n}\n",
	)
	p, results := testutil.LoadDirectory(t, directory)
	string_constant.Check(p, results)
	testutil.AssertBlocked(t, results, 1)
	testutil.AssertBlockedContains(
		t,
		results,
		`string literal "charlie delta" has constant alfa.Bravo`,
	)
}

func TestFlaggedFromParentConstantDirectory(t *testing.T) {
	directory := t.TempDir()
	testutil.WriteModFile(t, directory, "alfa")
	testutil.WriteFile(
		t,
		directory,
		"pkg/constant/constant.go",
		"package constant\n\nconst Bravo = \"charlie delta\"\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"pkg/echo/echo.go",
		"package echo\n\nfunc Echo() {\n\tprintln(\"charlie delta\")\n}\n",
	)
	p, results := testutil.LoadDirectory(
		t,
		filepath.Join(directory, "pkg/echo"),
	)
	string_constant.Check(p, results)
	testutil.AssertBlocked(t, results, 1)
	testutil.AssertBlockedContains(
		t,
		results,
		`string literal "charlie delta" has constant constant.Bravo`,
	)
}

func TestClean(t *testing.T) {
	temporary := testutil.PrepareTestPackage(t, "testdata/src/clean")
	testutil.WriteModFile(t, temporary, "clean")
	p, results := testutil.LoadDirectory(t, temporary)
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
	p, results := testutil.LoadDirectory(t, temporary)
	string_constant.Check(p, results)
	testutil.AssertBlocked(t, results, 5)
}

func TestTester(t *testing.T) {
	temporary := testutil.PrepareTestPackage(t, "testdata/src/tester")
	testutil.WriteModFile(t, temporary, "tester.test")
	p, results := testutil.LoadDirectory(t, temporary)
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
