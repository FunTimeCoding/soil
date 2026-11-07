package policy

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"mvdan.cc/sh/v3/syntax"
	"slices"
)

func (w *Walk) call(x *syntax.CallExpr) bool {
	result := true

	if len(x.Args) == 0 {
		for _, a := range x.Assigns {
			if !w.assign(a) {
				result = false
			}
		}

		return result
	}

	name, _ := literal(x.Args[0])

	for _, a := range x.Assigns {
		if a.Name == nil {
			result = w.block(constant.BlockerAssignment)

			continue
		}

		if a.Value != nil && !w.nested(a.Value) {
			result = false
		}

		locale := slices.Contains(constant.LocaleVariables, a.Name.Value)
		separator := a.Name.Value == constant.FieldSeparator &&
			name == constant.ReadCommand

		if !locale && !separator {
			result = w.block(constant.BlockerAssignment)
		}
	}

	for _, a := range x.Args {
		if !w.nested(a) {
			result = false
		}
	}

	w.commands++

	return w.invoke(x.Args) && result
}
