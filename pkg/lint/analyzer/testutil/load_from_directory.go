package testutil

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

func LoadFromDirectory(
	t *testing.T,
	directory string,
) (*packages.Package, *output.Results) {
	t.Helper()
	fileSet := token.NewFileSet()
	configuration := &packages.Config{
		Mode:  packages.LoadSyntax | packages.NeedModule,
		Fset:  fileSet,
		Dir:   directory,
		Tests: false,
	}
	loaded, e := packages.Load(configuration, constant.CurrentDirectory)

	if e != nil {
		t.Fatalf("load: %s", e)
	}

	if len(loaded) == 0 {
		t.Fatal("no packages loaded")
	}

	p := loaded[0]

	if len(p.Errors) > 0 {
		t.Fatalf("package errors: %v", p.Errors)
	}

	return p, output.NewResultsWithDirectory(
		fmt.Sprintf("%s/", directory),
	)
}
