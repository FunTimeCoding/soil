package error_answer

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/error_answer"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestBranchesAnsweringAgainstTheTable(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	p.PkgPath = "example.test/pkg/tool/sample/model_context"
	error_answer.Check(p, results)
	testutil.AssertBlocked(t, results, 6)
	testutil.AssertBlockedAt(t, results, "predicate.go", 8)
	testutil.AssertBlockedAt(t, results, "tool.go", 22)
	testutil.AssertBlockedAt(t, results, "tool.go", 26)
	testutil.AssertBlockedAt(t, results, "tool.go", 39)
	testutil.AssertBlockedAt(t, results, "rest.go", 19)
	testutil.AssertBlockedAt(t, results, "rest.go", 28)
	testutil.AssertBlockedContains(
		t,
		results,
		"validation answers response.Fail, this branch answers captureFail",
	)
	testutil.AssertBlockedContains(
		t,
		results,
		"unreachable answers captureFail, this branch answers response.Fail",
	)
	testutil.AssertBlockedContains(
		t,
		results,
		"not_found answers 404 Error, this branch answers 400 Error",
	)
	testutil.AssertBlockedContains(
		t,
		results,
		"an uncaptured 503 answers only unavailable",
	)
	testutil.AssertBlockedContains(
		t,
		results,
		"isClientError hides not_found, validation from the answer",
	)
}

func TestServerPackageInScope(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	p.PkgPath = "example.test/pkg/tool/sample/server"
	error_answer.Check(p, results)
	testutil.AssertBlocked(t, results, 6)
}

func TestOutOfScope(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	error_answer.Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}
