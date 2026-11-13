package resolve

import (
	"fmt"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func LoadPackagesTagged(
	directory string,
	tags []string,
	overlay map[string][]byte,
	patterns ...string,
) ([]*packages.Package, *token.FileSet, error) {
	set := token.NewFileSet()
	c := &packages.Config{
		Mode:       packages.LoadSyntax | packages.NeedModule,
		Fset:       set,
		Dir:        directory,
		Tests:      true,
		BuildFlags: TagFlags(tags),
		Overlay:    overlay,
	}
	result, e := packages.Load(c, patterns...)

	if e != nil {
		return nil, nil, fmt.Errorf("load packages: %w", e)
	}

	return result, set, nil
}
