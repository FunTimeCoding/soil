package imports

import (
	"go/types"
	"golang.org/x/tools/go/packages"
)

func ImportedPackage(
	p *packages.Package,
	importPath string,
) *types.Package {
	for _, i := range p.Types.Imports() {
		if i.Path() == importPath {
			return i
		}
	}

	return nil
}
