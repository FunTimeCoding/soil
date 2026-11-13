package preposition_name

import (
	"fmt"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/preposition_name"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

func analyze(t *testing.T) *output.Results {
	t.Helper()
	directory := testutil.PrepareTestPackage(t, "testdata/src/example")
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

	loaded := resolve.PreferTestVariants(all)
	faces := face.New(loaded)
	result := output.NewResultsWithDirectory(fmt.Sprintf("%s/", directory))

	for _, p := range loaded {
		if len(p.Errors) > 0 {
			t.Fatalf("package errors: %v", p.Errors)
		}

		preposition_name.Check(p, result, faces)
	}

	return result
}
