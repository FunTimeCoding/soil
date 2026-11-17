package symbol

import (
	"go/types"
	"golang.org/x/tools/go/packages"
	"slices"
)

func holderPackages(
	all []*packages.Package,
	modules []string,
	wanted map[*types.Named]bool,
) []string {
	var result []string

	for _, p := range all {
		if p.TypesInfo == nil ||
			!inWorkspace(modules, p.PkgPath) ||
			slices.Contains(result, p.PkgPath) {
			continue
		}

		for _, v := range p.TypesInfo.Types {
			if v.Type != nil && mentionsNamed(v.Type, wanted) {
				result = append(result, p.PkgPath)

				break
			}
		}
	}

	slices.Sort(result)

	return result
}
