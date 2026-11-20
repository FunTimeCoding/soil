package tester_function

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/tester_function"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestTesterFunction(t *testing.T) {
	directory := testutil.PrepareTestPackage(t, "testdata/src/example")
	configuration := &packages.Config{
		Mode:  packages.LoadSyntax | packages.NeedModule,
		Fset:  token.NewFileSet(),
		Dir:   directory,
		Tests: true,
	}
	loaded, e := packages.Load(configuration, "./...")

	if e != nil {
		t.Fatalf("load: %s", e)
	}

	results := output.NewResultsWithDirectory(fmt.Sprintf("%s/", directory))

	for _, p := range resolve.PreferTestVariants(loaded) {
		if len(p.Errors) > 0 {
			t.Fatalf("package errors: %v", p.Errors)
		}

		tester_function.Check(p, results)
	}

	testutil.AssertBlocked(t, results, 3)
	testutil.AssertBlockedContains(t, results, "Build is one of 3 functions")
	testutil.AssertBlockedContains(t, results, "Seed is one of 3 functions")
	testutil.AssertBlockedContains(
		t,
		results,
		"NewFixture is one of 3 functions",
	)
}
