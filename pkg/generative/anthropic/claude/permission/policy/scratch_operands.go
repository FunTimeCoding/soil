package policy

import "mvdan.cc/sh/v3/syntax"

func (w *Walk) scratchOperands(
	arguments []*syntax.Word,
	minimum int,
) bool {
	list, _ := operands(arguments)

	if len(list) < minimum {
		return false
	}

	for _, o := range list {
		if !w.scratch(o) {
			return false
		}
	}

	return true
}
