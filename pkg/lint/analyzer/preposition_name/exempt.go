package preposition_name

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"go/types"
)

func exempt(
	o types.Object,
	faces *face.Set,
	word string,
) bool {
	if faces.Implements(o) {
		return true
	}

	switch v := o.(type) {
	case *types.Func:
		s := v.Type().(*types.Signature)

		return state(s) || word == constant.MomentSegment && momentAnswer(s)
	case *types.Var:
		if v.IsField() {
			return isBool(v.Type()) || moment(v)
		}

		if packageLevel(v) {
			return momentName(v.Name())
		}

		return bare(v.Name()) || moment(v)
	case *types.Const:
		if packageLevel(v) {
			return momentName(v.Name())
		}

		return bare(v.Name()) || momentName(v.Name())
	}

	return false
}
