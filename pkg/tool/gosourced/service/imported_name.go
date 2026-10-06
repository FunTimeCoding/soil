package service

import "golang.org/x/tools/go/packages"

func importedName(
	all []*packages.Package,
	packagePath string,
) string {
	for _, p := range all {
		if i, okay := p.Imports[packagePath]; okay && i.Types != nil {
			return i.Types.Name()
		}
	}

	return ""
}
