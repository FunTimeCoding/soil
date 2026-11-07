package policy

import "mvdan.cc/sh/v3/syntax"

func (w *Walk) forgetOperands(arguments []*syntax.Word) {
	list, _ := operands(arguments)

	for _, o := range list {
		if value, okay := literal(o); okay {
			delete(w.variables, value)
		}
	}
}
