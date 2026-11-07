package policy

import "mvdan.cc/sh/v3/syntax"

func (w *Walk) statements(list []*syntax.Stmt) bool {
	result := true

	for _, s := range list {
		if !w.statement(s) {
			result = false
		}
	}

	return result
}
