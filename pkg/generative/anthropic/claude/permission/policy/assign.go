package policy

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"mvdan.cc/sh/v3/syntax"
	"slices"
)

func (w *Walk) assign(a *syntax.Assign) bool {
	if a.Name == nil || a.Index != nil || a.Array != nil {
		return w.block(constant.BlockerArray)
	}

	if slices.Contains(constant.DangerousVariables, a.Name.Value) {
		return w.block(fmt.Sprintf(constant.BlockerVariable, a.Name.Value))
	}

	if a.Value != nil && !w.nested(a.Value) {
		return false
	}

	value, okay := "", false

	if a.Value != nil && !a.Append {
		value, okay = w.resolve(a.Value)
	}

	if okay {
		w.variables[a.Name.Value] = value
	} else {
		delete(w.variables, a.Name.Value)
	}

	return true
}
