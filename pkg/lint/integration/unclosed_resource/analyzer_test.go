package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unclosed_resource"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestFlagged(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/flagged")
	unclosed_resource.Check(
		p,
		results,
		unclosed_resource.NewSummaries([]*packages.Package{p}),
	)
	testutil.AssertBlocked(t, results, 3)
	testutil.AssertBlockedContains(t, results, "PanicClose(r.Body)")
}

func TestClean(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/clean")
	unclosed_resource.Check(
		p,
		results,
		unclosed_resource.NewSummaries([]*packages.Package{p}),
	)
	testutil.AssertBlocked(t, results, 0)
}

func TestArrangedByConstructor(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/arranged")
	unclosed_resource.Check(
		p,
		results,
		unclosed_resource.NewSummaries([]*packages.Package{p}),
	)
	testutil.AssertBlocked(t, results, 0)
}
