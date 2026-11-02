package resolve

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/load_failure"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func LoadPackagesThrough(
	directory string,
	overlay map[string][]byte,
	patterns ...string,
) ([]*packages.Package, *token.FileSet, error) {
	set := token.NewFileSet()
	c := &packages.Config{
		Mode:       packages.LoadSyntax | packages.NeedModule,
		Fset:       set,
		Dir:        directory,
		Tests:      true,
		BuildFlags: BuildFlags(directory),
		Overlay:    overlay,
	}
	result, e := packages.Load(c, patterns...)

	if e != nil {
		return nil, nil, fmt.Errorf("load packages: %w", e)
	}

	for _, p := range result {
		for _, f := range p.Errors {
			return nil, nil, load_failure.New(p.PkgPath, f)
		}
	}

	return result, set, nil
}
