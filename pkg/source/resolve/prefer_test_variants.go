package resolve

import (
	"golang.org/x/tools/go/packages"
	"strings"
)

func PreferTestVariants(loaded []*packages.Package) []*packages.Package {
	best := make(map[string]*packages.Package)
	var order []string

	for _, p := range loaded {
		if strings.HasSuffix(p.PkgPath, ".test") && p.Name == "main" {
			continue
		}

		current, exists := best[p.PkgPath]

		if !exists {
			best[p.PkgPath] = p
			order = append(order, p.PkgPath)

			continue
		}

		if len(p.Syntax) > len(current.Syntax) {
			best[p.PkgPath] = p
		}
	}

	var result []*packages.Package

	for _, path := range order {
		if len(best[path].Syntax) == 0 {
			continue
		}

		result = append(result, best[path])
	}

	return result
}
