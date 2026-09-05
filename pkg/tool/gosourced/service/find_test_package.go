package service

import "golang.org/x/tools/go/packages"

func findTestPackage(
	all []*packages.Package,
	packagePath string,
) *packages.Package {
	variant := testVariantPath(packagePath)
	seen := make(map[string]bool)
	queue := append([]*packages.Package{}, all...)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if seen[p.ID] {
			continue
		}

		seen[p.ID] = true

		if len(p.GoFiles) > 0 &&
			(p.PkgPath == packagePath || p.PkgPath == variant) {
			return p
		}

		for _, i := range p.Imports {
			queue = append(queue, i)
		}
	}

	return nil
}
