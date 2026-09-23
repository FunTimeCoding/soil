package resolve

import (
	"fmt"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func LoadPackages(
	directory string,
	patterns ...string,
) ([]*packages.Package, *token.FileSet, error) {
	set := token.NewFileSet()
	c := &packages.Config{
		Mode:       packages.LoadSyntax | packages.NeedModule,
		Fset:       set,
		Dir:        directory,
		Tests:      true,
		BuildFlags: BuildFlags(directory),
	}
	result, e := packages.Load(c, patterns...)

	if e != nil {
		return nil, nil, fmt.Errorf("load packages: %w", e)
	}

	for _, p := range result {
		for _, f := range p.Errors {
			return nil, nil, fmt.Errorf("%s: %w", p.PkgPath, f)
		}
	}

	return result, set, nil
}
