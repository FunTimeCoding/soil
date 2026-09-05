package omit_empty_zero

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/omit_empty_zero"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestOmitEmptyZero(t *testing.T) {
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

		omit_empty_zero.Check(p, results)
	}

	testutil.AssertBlocked(t, results, 4)
	testutil.AssertBlockedContains(t, results, "field Count")
	testutil.AssertBlockedContains(t, results, "field Depth")
	testutil.AssertBlockedContains(t, results, "use omitzero")
}
