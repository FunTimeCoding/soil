package policy

import "mvdan.cc/sh/v3/syntax"

func (w *Walk) declare(x *syntax.DeclClause) bool {
	result := true

	for _, a := range x.Args {
		if !w.assign(a) {
			result = false
		}
	}

	return result
}
