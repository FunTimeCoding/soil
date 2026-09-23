package tester_receiver

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/tester_receiver"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestTesterReceiver(t *testing.T) {
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

		tester_receiver.Check(p, results)
	}

	testutil.AssertBlocked(t, results, 2)
	testutil.AssertBlockedContains(t, results, "Bad takes *testing.T")
	testutil.AssertBlockedContains(t, results, "AlsoBad takes *testing.T")
}
