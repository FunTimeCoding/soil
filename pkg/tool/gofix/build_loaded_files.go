package gofix

import "golang.org/x/tools/go/packages"

func BuildLoadedFiles(all []*packages.Package) map[string]bool {
	result := make(map[string]bool)

	for _, p := range all {
		for _, f := range p.GoFiles {
			result[f] = true
		}
	}

	return result
}
