package imports

import "golang.org/x/tools/go/packages"

func ModulePath(
	p *packages.Package,
	importPath string,
) string {
	imported, okay := p.Imports[importPath]

	if !okay || imported.Module == nil {
		return ""
	}

	return imported.Module.Path
}
