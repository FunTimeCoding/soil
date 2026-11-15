package testutil

import (
	"fmt"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

func loadVariants(
	t *testing.T,
	directory string,
) ([]*packages.Package, *output.Results) {
	t.Helper()
	configuration := &packages.Config{
		Mode:  packages.LoadSyntax | packages.NeedModule,
		Fset:  token.NewFileSet(),
		Dir:   directory,
		Tests: true,
	}
	all, e := packages.Load(configuration, goModConstant.AllPackages)

	if e != nil {
		t.Fatalf("load: %s", e)
	}

	result := resolve.PreferTestVariants(all)

	for _, p := range result {
		if len(p.Errors) > 0 {
			t.Fatalf("package errors: %v", p.Errors)
		}
	}

	return result, output.NewResultsWithDirectory(fmt.Sprintf("%s/", directory))
}
