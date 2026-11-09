package unguarded_peer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/types"
)

func isPeerClient(t types.Type) bool {
	i, okay := t.Underlying().(*types.Interface)

	if !okay {
		return false
	}

	for m := range i.Methods() {
		if m.Name() == constant.ReachableMethod {
			return true
		}
	}

	return false
}
