package preposition_name

import "go/types"

func packageLevel(o types.Object) bool {
	return o.Pkg() != nil && o.Parent() == o.Pkg().Scope()
}
