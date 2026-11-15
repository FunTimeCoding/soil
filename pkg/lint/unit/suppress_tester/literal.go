package suppress_tester

import "go/ast"

func (o *Tester) Literal(line int) *ast.CompositeLit {
	o.t.Helper()

	return firstOnLine[*ast.CompositeLit](o, line)
}
