package face

import (
	"github.com/funtimecoding/soil/pkg/source/index"
	"go/types"
)

func methodSignatures(named *types.Named) map[string]string {
	set := types.NewMethodSet(types.NewPointer(named))
	result := make(map[string]string, set.Len())

	for i := range set.Len() {
		if f, okay := set.At(i).Obj().(*types.Func); okay {
			result[f.Name()] = index.Signature(f)
		}
	}

	return result
}
