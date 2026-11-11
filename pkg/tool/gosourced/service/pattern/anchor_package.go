package pattern

import "go/types"

func anchorPackage(
	declaration types.Object,
	receiver string,
) string {
	if receiver != "" || declaration.Pkg() == nil ||
		declaration.Parent() != declaration.Pkg().Scope() {
		return ""
	}

	if _, isType := declaration.(*types.TypeName); isType {
		return ""
	}

	return declaration.Pkg().Path()
}
