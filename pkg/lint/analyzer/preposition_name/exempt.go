package preposition_name

import (
	"github.com/funtimecoding/soil/pkg/lint/face"
	"go/types"
)

func exempt(
	o types.Object,
	faces *face.Set,
) bool {
	if faces.Implements(o) {
		return true
	}

	switch v := o.(type) {
	case *types.Func:
		s := v.Type().(*types.Signature)

		return state(s) || momentAnswer(s)
	case *types.Var:
		return v.IsField() && (isBool(v.Type()) || moment(v))
	}

	return false
}
