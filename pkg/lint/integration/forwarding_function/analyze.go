package forwarding_function

import (
	"fmt"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

func analyze(
	t *testing.T,
	check func(*packages.Package, *output.Results),
) *output.Results {
	t.Helper()
	directory := testutil.PrepareTestPackage(t, "testdata/src/example")
	configuration := &packages.Config{
		Mode:  packages.LoadSyntax | packages.NeedModule,
		Fset:  token.NewFileSet(),
		Dir:   directory,
		Tests: true,
	}
	loaded, e := packages.Load(configuration, goModConstant.AllPackages)

	if e != nil {
		t.Fatalf("load: %s", e)
	}

	result := output.NewResultsWithDirectory(fmt.Sprintf("%s/", directory))

	for _, p := range resolve.PreferTestVariants(loaded) {
		if len(p.Errors) > 0 {
			t.Fatalf("package errors: %v", p.Errors)
		}

		check(p, result)
	}

	return result
}
