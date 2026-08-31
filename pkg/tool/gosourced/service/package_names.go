package service

import "golang.org/x/tools/go/packages"

func packageNames(all []*packages.Package) map[string][]string {
	result := map[string][]string{}
	seen := map[string]bool{}
	queue := append([]*packages.Package{}, all...)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if seen[p.ID] {
			continue
		}

		seen[p.ID] = true

		if p.Types != nil {
			name := p.Types.Name()
			known := false

			for _, path := range result[name] {
				if path == p.PkgPath {
					known = true
				}
			}

			if !known && name != "" && name != "main" {
				result[name] = append(result[name], p.PkgPath)
			}
		}

		for _, imported := range p.Imports {
			queue = append(queue, imported)
		}
	}

	return result
}
