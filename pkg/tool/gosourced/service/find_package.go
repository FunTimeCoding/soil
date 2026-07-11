package service

import "golang.org/x/tools/go/packages"

func findPackage(
	all []*packages.Package,
	packagePath string,
) *packages.Package {
	seen := make(map[string]bool)
	queue := append([]*packages.Package{}, all...)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if seen[p.PkgPath] {
			continue
		}

		seen[p.PkgPath] = true

		if p.PkgPath == packagePath {
			return p
		}

		for _, i := range p.Imports {
			queue = append(queue, i)
		}
	}

	return nil
}
