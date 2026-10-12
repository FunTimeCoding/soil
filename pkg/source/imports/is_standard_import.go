package imports

import "golang.org/x/tools/go/packages"

func IsStandardImport(
	p *packages.Package,
	importPath string,
) bool {
	imported, okay := p.Imports[importPath]

	return okay && imported.Module == nil
}
