package resolve

import "golang.org/x/tools/go/packages"

func PackagePaths(listed []*packages.Package) map[string]bool {
	result := make(map[string]bool)

	for _, p := range listed {
		result[p.PkgPath] = true
	}

	return result
}
