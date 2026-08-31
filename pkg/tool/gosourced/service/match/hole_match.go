package match

import "go/ast"

func (m *Matcher) holeMatch(name string, site ast.Node) bool {
	expression, okay := site.(ast.Expr)

	if !okay {
		return false
	}

	if previous, bound := m.bindings[name]; bound {
		if m.render(previous) != m.render(expression) {
			return false
		}
	}

	if !m.assignable(expression, m.holes[name]) {
		return false
	}

	m.bindings[name] = expression

	return true
}
