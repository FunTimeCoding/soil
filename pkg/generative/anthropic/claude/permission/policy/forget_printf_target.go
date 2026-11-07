package policy

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"mvdan.cc/sh/v3/syntax"
)

func (w *Walk) forgetPrintfTarget(arguments []*syntax.Word) {
	for i, a := range arguments {
		value, okay := literal(a)

		if okay && value == constant.VariableOption && i+1 < len(arguments) {
			if name, named := literal(arguments[i+1]); named {
				delete(w.variables, name)
			}
		}
	}
}
