package suppress_tester

import "go/ast"

func (o *Tester) Call(line int) *ast.CallExpr {
	o.t.Helper()

	return firstOnLine[*ast.CallExpr](o, line)
}
