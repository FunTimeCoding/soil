package symbol

import (
	"golang.org/x/tools/go/packages"
	"slices"
)

func workspaceModules(all []*packages.Package) []string {
	var result []string

	for _, p := range all {
		m := p.Module

		if m == nil || !m.Main && m.Replace == nil {
			continue
		}

		if !slices.Contains(result, m.Path) {
			result = append(result, m.Path)
		}
	}

	return result
}
