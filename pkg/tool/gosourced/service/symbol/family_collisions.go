package symbol

import (
	"fmt"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func familyCollisions(
	all []*packages.Package,
	members []*types.Func,
	newName string,
) []string {
	joined := map[types.Object]bool{}

	for _, f := range members {
		joined[f] = true
	}

	var result []string

	for _, t := range namedTypes(all) {
		var holder types.Type = types.NewPointer(t)

		if _, isInterface := t.Underlying().(*types.Interface); isInterface {
			holder = t
		}

		o, _, _ := types.LookupFieldOrMethod(
			holder,
			true,
			members[0].Pkg(),
			members[0].Name(),
		)

		if !joined[o] {
			continue
		}

		taken, _, _ := types.LookupFieldOrMethod(
			holder,
			true,
			members[0].Pkg(),
			newName,
		)

		if taken == nil {
			continue
		}

		result = append(
			result,
			fmt.Sprintf("%s already has %s", types.TypeString(t, nil), newName),
		)
	}

	return result
}
