package resolve

import "golang.org/x/tools/go/packages"

func NewNames(all []*packages.Package) *Names {
	names := make(map[string]string)
	var visit func(p *packages.Package)
	visit = func(p *packages.Package) {
		if _, seen := names[p.PkgPath]; seen || p.Name == "" {
			return
		}

		names[p.PkgPath] = p.Name

		for _, imported := range p.Imports {
			visit(imported)
		}
	}

	for _, p := range all {
		visit(p)
	}

	return &Names{names: names}
}
