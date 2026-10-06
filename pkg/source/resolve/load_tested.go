package resolve

import (
	"fmt"
	"golang.org/x/tools/go/packages"
)

func LoadTested(
	directory string,
	patterns ...string,
) ([]*packages.Package, error) {
	c := &packages.Config{
		Mode:       packages.LoadSyntax | packages.NeedModule,
		Dir:        directory,
		Tests:      true,
		BuildFlags: BuildFlags(directory),
	}
	result, e := packages.Load(c, patterns...)

	if e != nil {
		return nil, fmt.Errorf("load packages: %w", e)
	}

	return PreferTestVariants(result), nil
}
