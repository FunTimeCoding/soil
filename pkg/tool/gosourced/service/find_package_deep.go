package service

import "golang.org/x/tools/go/packages"

func findPackageDeep(
	all []*packages.Package,
	packagePath string,
) *packages.Package {
	seen := map[string]bool{}
	queue := append([]*packages.Package{}, all...)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if seen[p.ID] {
			continue
		}

		seen[p.ID] = true

		if p.PkgPath == packagePath && p.Types != nil {
			return p
		}

		for _, imported := range p.Imports {
			queue = append(queue, imported)
		}
	}

	return nil
}
