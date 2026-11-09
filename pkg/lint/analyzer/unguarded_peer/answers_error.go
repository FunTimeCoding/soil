package unguarded_peer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/types"
)

func answersError(o types.Object) bool {
	r := o.Type().(*types.Signature).Results()
	failure := types.Universe.Lookup(constant.ErrorType).Type()

	for i := range r.Len() {
		if types.Identical(r.At(i).Type(), failure) {
			return true
		}
	}

	return false
}
