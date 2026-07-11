package service

import "golang.org/x/tools/go/packages"

func importsTransitively(
	from *packages.Package,
	packagePath string,
) bool {
	seen := make(map[string]bool)
	var queue []*packages.Package

	for _, i := range from.Imports {
		queue = append(queue, i)
	}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if seen[p.PkgPath] {
			continue
		}

		seen[p.PkgPath] = true

		if p.PkgPath == packagePath {
			return true
		}

		for _, i := range p.Imports {
			queue = append(queue, i)
		}
	}

	return false
}
