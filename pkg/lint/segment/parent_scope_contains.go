package segment

import "go/types"

func ParentScopeContains(
	scope *types.Scope,
	name string,
) bool {
	for p := scope.Parent(); p != nil; p = p.Parent() {
		if p.Parent() == nil {
			return false
		}

		if p.Lookup(name) != nil {
			return true
		}
	}

	return false
}
