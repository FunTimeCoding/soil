package module

import "go/types"

func objectBelongsToModule(
	o types.Object,
	modulePath string,
) bool {
	return o != nil &&
		o.Pkg() != nil &&
		belongsToModule(o.Pkg().Path(), modulePath)
}
